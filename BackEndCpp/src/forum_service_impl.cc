#include "../proto/forum.pb.h"
// #include "../proto/forum.grpc.pb.h"
#include <grpcpp/grpcpp.h>
#include <iostream>
#include <string>
#include <ctime>
#include <iomanip>
#include <sstream>
#include <sodium.h>
// #include "../service/user.hpp"
#include "forum_service_impl.hpp"


// class ForumServiceImpl final : public proto::ForumService::Service {
//  public:
//     explicit ForumServiceImpl(sql::Connection* dbConnection) : con(dbConnection) {}
//     grpc::Status CreateUser(grpc::ServerContext* context, const proto::CreateUserRequest* request, proto::CreateUserResponse* response) override {
//         auto defaultRoleId = 1;
//         std::cout << "CreateUser called" << std::endl;
//         std::tm tm = {};
//         std::istringstream ss(request->userdatebirth());
//         ss >> std::get_time(&tm, "%Y/%m/%d");
//         if (ss.fail()) {
//             return grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, "Invalid date format, should be YYYY/MM/DD");
//         }
//         auto dateOfBirth = std::mktime(&tm);

//         if (request->userpassword() != request->userrepassword()) {
//             return grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, "Passwords do not match");
//         }

//         std::string hashedPword;
//         // hash password
//         try {
//             // hashedPword = BCrypt::generateHash(request->userpassword());
//             unsigned char hash[crypto_generichash_BYTES];
//             auto hashedPword = crypto_generichash(hash, sizeof(hash), reinterpret_cast<const unsigned char*>(request->userpassword().c_str()), request->userpassword().size(), nullptr, 0);
//         } catch (const std::exception& e) {
//             return grpc::Status(grpc::StatusCode::INTERNAL, "Failed to hash password");
//         }
//         User usr(request->username(), hashedPword, request->usersex(), request->userage(), defaultRoleId, request->useremail(), tm, std::nullopt);
//         usr.saveToDatabase(con, usr);

//         return grpc::Status::OK;
//     }
//  private:
//     // Database db;
//     sql::Connection* con;
// };

ForumServiceImpl::ForumServiceImpl(sql::Connection* dbConnection) : con(dbConnection) {}

grpc::Status ForumServiceImpl::CreateUser(grpc::ServerContext* context, const proto::CreateUserRequest* request, proto::CreateUserResponse* response) {
    auto defaultRoleId = 1;
    std::cout << "CreateUser called" << std::endl;
    std::tm tm = {};
    std::istringstream ss(request->userdatebirth());
    ss >> std::get_time(&tm, "%Y/%m/%d");
    if (ss.fail()) {
        return grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, "Invalid date format, should be YYYY/MM/DD");
    }
    auto dateOfBirth = std::mktime(&tm);

    if (request->userpassword() != request->userrepassword()) {
        return grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, "Passwords do not match");
    }

    std::string hashedPword;
    // hash password
    try {
        // hashedPword = BCrypt::generateHash(request->userpassword());
        unsigned char hash[crypto_generichash_BYTES];
        auto hashedPword = crypto_generichash(hash, sizeof(hash), reinterpret_cast<const unsigned char*>(request->userpassword().c_str()), request->userpassword().size(), nullptr, 0);
    } catch (const std::exception& e) {
        return grpc::Status(grpc::StatusCode::INTERNAL, "Failed to hash password");
    }
    User usr(request->username(), hashedPword, request->usersex(), request->userage(), defaultRoleId, request->useremail(), tm, std::nullopt);
    usr.saveToDatabase(con, usr);
    return grpc::Status::OK;
}