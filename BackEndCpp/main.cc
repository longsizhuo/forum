#include <iostream>
#include <fstream>
#include <string>
#include <mysql/mysql.h>
#include <grpcpp/grpcpp.h>
#include "proto/forum.grpc.pb.h"
#include <yaml-cpp/yaml.h>

class Database {
public:
    MYSQL *conn;

    Database(const std::string &dsn) {
        conn = mysql_init(nullptr);
        if (conn == nullptr) {
            throw std::runtime_error("mysql_init() failed");
        }

        if (mysql_real_connect(conn, "host", "user", "password", "database", 0, nullptr, 0) == nullptr) {
            mysql_close(conn);
            throw std::runtime_error("mysql_real_connect() failed");
        }
    }

    ~Database() {
        if (conn != nullptr) {
            mysql_close(conn);
        }
    }

    void AutoMigrate() {
        // 假设这里实现了数据表的迁移
        std::cout << "AutoMigrate executed." << std::endl;
    }
};

void InitConfig(YAML::Node &config) {
    std::string basePath = std::getenv("BASE_PATH");
    if (basePath.empty()) {
        basePath = "./BackEnd";
    }

    std::string configPath = basePath + "/config/app.yml";
    std::cout << "configPath: " << configPath << std::endl;

    try {
        config = YAML::LoadFile(configPath);
        std::cout << "Config app: " << config["app"] << std::endl;
        std::cout << "Config mysql: " << config["mysql"] << std::endl;
    } catch (const std::exception &e) {
        std::cerr << "Failed to read config file: " << e.what() << std::endl;
        throw;
    }
}