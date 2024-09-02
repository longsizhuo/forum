#include <string>
#include <ctime>
#include <optional>
#include <iostream>
#include <mysql/mysql.h>
#include <mysql_connection.h>
#include <cppconn/prepared_statement.h>
#include <mysql_driver.h>
#include <cppconn/statement.h>
#include <cppconn/resultset.h>
#include <cppconn/exception.h>

class User {
public:
    // 构造函数   
    User() = default;
    User(const std::string& userName,
         const std::string& userPassword,
         const std::string& userSex,
         int userAge,
         int roleID,
         const std::string& userEmail,
         const std::tm& dateBirth,
         std::optional<int> uid = std::nullopt);

    // Getter 方法
    int getUID() const;
    std::string getUserName() const;
    std::string getUserPassword() const;
    std::string getUserSex() const;
    int getUserAge() const;
    int getRoleID() const;
    std::string getUserEmail() const;
    std::tm getDateBirth() const;

    // Setter 方法
    void setUID(int uid);
    void setUserName(const std::string& userName);
    void setUserPassword(const std::string& userPassword);
    void setUserSex(const std::string& userSex);
    void setUserAge(int userAge);
    void setRoleID(int roleID);
    void setUserEmail(const std::string& userEmail);
    void setDateBirth(const std::tm& dateBirth);

    // 数据库操作方法
    bool saveToDatabase(sql::Connection* con, User usr) const;
    static std::optional<User> getUserByEmail(sql::Connection& con, const std::string& email);
    static std::optional<User> getUserByUID(sql::Connection& con, int uid);

    // 验证方法
    static bool isEmailRegistered(sql::Connection& con, const std::string& email);
    bool validatePassword(const std::string& password) const;

    // 辅助方法
    std::string toString() const;

private:
    std::optional<int> uid_;
    std::string userName_;
    std::string userPassword_;
    std::string userSex_;
    int userAge_;
    int roleID_;
    std::string userEmail_;
    std::tm dateBirth_;
};

class Database {
// public:
//     MYSQL *conn;
//     sql::Connection* con;
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
public:
    sql::Connection* con;

    Database() = default;

    Database(const std::string& host, const std::string& user, const std::string& password, const std::string& dbname) {
        try {
            sql::mysql::MySQL_Driver* driver = sql::mysql::get_mysql_driver_instance();
            con = driver->connect("tcp://" + host + ":3306", user, password);
            con->setSchema(dbname);
            std::cout << "Database connected." << std::endl;
        } catch (sql::SQLException& e) {
            std::cerr << "Error connecting to database: " << e.what() << std::endl;
            throw;
        }
    }

    ~Database() {
        if (con != nullptr) {
            delete(con);
        }
    }

    // void executeSQL(const std::string &sql) {
    //     if (mysql_query(con, sql.c_str())) {
    //         throw std::runtime_error("MySQL query failed: " + std::string(mysql_error(conn)));
    //     }
    // }
    void executeSQL(const std::string& sql) {
        try {
            sql::Statement* stmt = con->createStatement();
            stmt->execute(sql);
            delete stmt;
        } catch (sql::SQLException& e) {
            std::cerr << "Error executing SQL: " << e.what() << std::endl;
            throw;
        }
    }


    private:
        std::string host_;
        std::string user_;
        std::string password_;
        std::string dbname_;
};