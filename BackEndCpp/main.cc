#include <iostream>
// #include <fstream>
#include <string>
#include <mysql/mysql.h>
#include <grpcpp/grpcpp.h>
// #include "proto/forum.grpc.pb.h"
#include <yaml-cpp/yaml.h>
#include <regex>
#include "mysqlDb.hpp"

using namespace sqlDb;
class Database {
public:
    MYSQL *conn;

    Database(const std::string &host, const std::string &user, const std::string &password, const std::string &dbname) {
        conn = mysql_init(nullptr);
        if (conn == nullptr) {
            throw std::runtime_error("mysql_init() failed");
        }

        if (mysql_real_connect(conn, host.c_str(), user.c_str(), password.c_str(), dbname.c_str(), 0, nullptr, 0) == nullptr) {
            mysql_close(conn);
            throw std::runtime_error("mysql_real_connect() failed");
        }
        std::cout << "Database connected." << std::endl;
    }

    ~Database() {
        if (conn != nullptr) {
            mysql_close(conn);
        }
    }

    void executeSQL(const std::string &sql) {
        if (mysql_query(conn, sql.c_str())) {
            throw std::runtime_error("MySQL query failed: " + std::string(mysql_error(conn)));
        }
    }

    void AutoMigrate() {
        // 假设这里实现了数据表的迁移
        sqlDb::AuthRole authRole;
        executeSQL(authRole.createAuthRoleTable.c_str());
        sqlDb::User user;
        executeSQL(user.createUserTable.c_str());
        sqlDb::Admin admin;
        executeSQL(admin.createAdminTable.c_str());
        sqlDb::Section section;
        executeSQL(section.createSectionTable.c_str());
        sqlDb::Topic topic;
        executeSQL(topic.createTopicTable.c_str());
        sqlDb::Reply reply;
        executeSQL(reply.createReplyTable.c_str());
        sqlDb::Channel channel;
        executeSQL(channel.createChannelTable.c_str());
        sqlDb::Message message;
        executeSQL(message.createMessageTable.c_str());
        sqlDb::Media media;
        executeSQL(media.createMediaTable.c_str());


        std::cout << "AutoMigrate executed." << std::endl;
    }
};

void InitConfig(YAML::Node &config) {
    const char* basePathEnv = std::getenv("BASE_PATH");
    std::string basePath;
    
    if (basePathEnv == nullptr || std::string(basePathEnv).empty()) {
        basePath = "./BackEndCpp";
    } else {
        basePath = basePathEnv;
    }

    std::string configPath = basePath + "/config/app.yml";
    std::cout << "configPath: " << configPath << std::endl;

    try {
        config = YAML::LoadFile(configPath);
        // std::cout << "Config content: " << config << std::endl;
        std::cout << "Config mysql: " << config["mysql"] << std::endl;
    } catch (const std::exception &e) {
        std::cerr << "Failed to read config file: " << e.what() << std::endl;
        throw;
    }
}

void ParseDNS(const std::string &dns, std::string &host, std::string &user, std::string &password, std::string &dbname) {
    std::regex rgx(R"(([^:]+):([^@]+)@tcp\(([^:]+):\d+\)/([^?]+))");
    std::smatch match;
    auto cnt = 0;
    // std::cout << "dns: " << dns << std::endl;
    std::cout << match.size() << std::endl;
    for (const auto &m : match) {
        std::cout << "match[" << cnt << "]: ";
        std::cout << m << std::endl;
        ++cnt;
    }
    if (std::regex_search(dns, match, rgx)) {
        user = match[1];
        password = match[2];
        host = match[3];
        dbname = match[4];
    } else {
        throw std::runtime_error("Failed to parse DNS string.");
    }
}

int main() {
    YAML::Node config;

    try {
        InitConfig(config);
    } catch (const std::exception &e) {
        std::cerr << "Initialization failed: " << e.what() << std::endl;
        return EXIT_FAILURE;
    }

    try {
        std::string dns = config["mysql"]["dns"].as<std::string>();

        std::string host, user, password, dbname;
        ParseDNS(dns, host, user, password, dbname);
        std::cout << "host: " << host << ", user: " << user << ", password: " << password << ", dbname: " << dbname << std::endl;
        Database db(host, user, password, dbname);
        db.AutoMigrate();
    } catch (const std::exception &e) {
        std::cerr << "Database connection failed: " << e.what() << std::endl;
        return EXIT_FAILURE;
    }

    std::cout << "Program finished successfully." << std::endl;
    return EXIT_SUCCESS;
}