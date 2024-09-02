// #include <iostream>
// #include <string>
// #include <mysql/mysql.h>
// #include <grpcpp/grpcpp.h>

// class Database {
// public:
//     MYSQL *conn;
//     Database() = default;
//     Database(const std::string &host, const std::string &user, const std::string &password, const std::string &dbname) 
//     : host_(host), user_(user), password_(password), dbname_(dbname) {
//         conn = mysql_init(nullptr);
//         if (conn == nullptr) {
//             throw std::runtime_error("mysql_init() failed");
//         }

//         if (mysql_real_connect(conn, host.c_str(), user.c_str(), password.c_str(), dbname.c_str(), 0, nullptr, 0) == nullptr) {
//             mysql_close(conn);
//             throw std::runtime_error("mysql_real_connect() failed");
//         }
//         std::cout << "Database connected." << std::endl;
//     }

//     ~Database() {
//         if (conn != nullptr) {
//             mysql_close(conn);
//         }
//     }

//     void executeSQL(const std::string &sql) {
//         if (mysql_query(conn, sql.c_str())) {
//             throw std::runtime_error("MySQL query failed: " + std::string(mysql_error(conn)));
//         }
//     }

//     private:
//         std::string host_;
//         std::string user_;
//         std::string password_;
//         std::string dbname_;
// };
