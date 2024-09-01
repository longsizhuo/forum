#include "../proto/forum.pb.h"
#include "../proto/forum.grpc.pb.h"
#include <grpcpp/grpcpp.h>
#include <iostream>
#include <memory>
#include <string>
#include <ctime>
#include <iomanip>
#include <sstream>
#include <bcrypt/BCrypt.hpp>
#include "service/database.hpp"


class ForumServiceImpl final : public proto::ForumService::Service {
 public:
    grpc::Status CreateUser(grpc::ServerContext* context, const proto::CreateUserRequest* request, proto::CreateUserResponse* response) override {
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
            hashedPword = BCrypt::generateHash(request->userpassword());
        } catch (const std::exception& e) {
            return grpc::Status(grpc::StatusCode::INTERNAL, "Failed to hash password");
        }

        //TODO: check email is unique
        return grpc::Status::OK;
    }
 private:
    Database db;
};