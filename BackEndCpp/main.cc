#include <iostream>
#include <memory>
#include <grpcpp/grpcpp.h>
#include <mysql/mysql.h>
#include <mysql_connection.h>
#include <cppconn/prepared_statement.h>
#include "src/forum_service_impl.hpp"
#include <yaml-cpp/yaml.h>
#include <regex>
#include <string>
#include <cstdlib>
#include <ctime>
#include <iomanip>



// using Database = User::Database;
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
        // std::cout << “Config content: ” << config << std::endl;
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
    // std::cout << “dns: ” << dns << std::endl;
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
        // 创建ForumServiceImpl实例并传递数据库连接
        ForumServiceImpl service(db.con);
        // 启动gRPC服务器并注册服务
        grpc::ServerBuilder builder;
        builder.RegisterService(&service);
        std::unique_ptr<grpc::Server> server(builder.BuildAndStart());
        std::cout << "Server started" << std::endl;
        server->Wait();
        User usr;
        usr.setUserName("test");
        usr.setUserPassword("test");
        usr.setUserSex("test");
        usr.setUserAge(1);
        usr.setRoleID(1);
        usr.setUserEmail("test");
        std::tm tm = {};
        std::istringstream ss("2021/01/01");
        ss >> std::get_time(&tm, "%Y/%m/%d");
        usr.setDateBirth(tm);
        usr.saveToDatabase(db.con, usr);
        
        return 0;
    } catch (const std::exception &e) {
        std::cerr << "Database connection failed: " << e.what() << std::endl;
        return EXIT_FAILURE;
    }
    

}