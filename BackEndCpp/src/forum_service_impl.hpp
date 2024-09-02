#ifndef FORUM_SERVICE_IMPL_H
#define FORUM_SERVICE_IMPL_H

#include "../proto/forum.grpc.pb.h"
#include <grpcpp/grpcpp.h>
#include <mysql_connection.h>
#include "../service/user.hpp"  // 包含用户相关的头文件

class ForumServiceImpl final : public proto::ForumService::Service {
public:
    explicit ForumServiceImpl(sql::Connection* dbConnection);
    grpc::Status CreateUser(grpc::ServerContext* context, const proto::CreateUserRequest* request, proto::CreateUserResponse* response) override;

private:
    sql::Connection* con;
};

#endif // FORUM_SERVICE_IMPL_H