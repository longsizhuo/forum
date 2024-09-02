#include "user.hpp"
#include <cppconn/exception.h>
#include <memory>
#include <sodium.h>
#include <sstream>
#include <iomanip>
#include <string>
#include <ctime>
#include "database.hpp"
#include <mysql_connection.h>
#include <cppconn/prepared_statement.h>
#include <optional>

// 构造函数实现
User::User(const std::string& userName,
           const std::string& userPassword,
           const std::string& userSex,
           int userAge,
           int roleID,
           const std::string& userEmail,
           const std::tm& dateBirth,
           std::optional<int> uid)
    : userName_(userName),
      userPassword_(userPassword),
      userSex_(userSex),
      userAge_(userAge),
      roleID_(roleID),
      userEmail_(userEmail),
      dateBirth_(dateBirth),
      uid_(uid) {}

// Getter 方法实现
int User::getUID() const {
    return uid_.value_or(-1);
}

std::string User::getUserName() const {
    return userName_;
}

std::string User::getUserPassword() const {
    return userPassword_;
}

std::string User::getUserSex() const {
    return userSex_;
}

int User::getUserAge() const {
    return userAge_;
}

int User::getRoleID() const {
    return roleID_;
}

std::string User::getUserEmail() const {
    return userEmail_;
}

std::tm User::getDateBirth() const {
    return dateBirth_;
}

// Setter 方法实现
void User::setUID(int uid) {
    uid_ = uid;
}

void User::setUserName(const std::string& userName) {
    userName_ = userName;
}

void User::setUserPassword(const std::string& userPassword) {
    // unsigned char hash[crypto_generichash_BYTES];
    // auto hashed = crypto_generichash(hash, sizeof(hash), reinterpret_cast<const unsigned char*>(userPassword.c_str()), userPassword.size(), nullptr, 0);
    userPassword_ = userPassword;
}

void User::setUserSex(const std::string& userSex) {
    userSex_ = userSex;
}

void User::setUserAge(int userAge) {
    userAge_ = userAge;
}

void User::setRoleID(int roleID) {
    roleID_ = roleID;
}

void User::setUserEmail(const std::string& userEmail) {
    userEmail_ = userEmail;
}

void User::setDateBirth(const std::tm& dateBirth) {
    dateBirth_ = dateBirth;
}
// TODO: Database operation methods
bool User::saveToDatabase(sql::Connection* con, User usr) const {
    try {
    std::unique_ptr<sql::PreparedStatement> statement(con->prepareStatement(
    "INSERT INTO User (userName, userPassword, userSex, userAge, roleID, userEmail, dateBirth, uid) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"));
    statement->setString(1, usr.getUserName());
    statement->setString(2, usr.getUserPassword());
    statement->setString(3, usr.getUserSex());
    statement->setInt(4, usr.getUserAge());
    statement->setInt(5, usr.getRoleID());
    statement->setString(6, usr.getUserEmail());
    std::ostringstream oss;
        oss << std::put_time(&usr.dateBirth_, "%Y-%m-%d %H:%M:%S"); // 采用标准的时间格式
        statement->setString(7, oss.str());

        if (usr.uid_.has_value()) {
            statement->setInt(8, usr.uid_.value());
        } else {
            statement->setNull(8, sql::DataType::INTEGER);
        }

        statement->executeUpdate();
        return true;
    } catch (sql::SQLException& e) {
        std::cerr << "SQL error: " << e.what() << std::endl;
        return false;
    }
}

std::optional<User> User::getUserByEmail(sql::Connection& con, const std::string& email) {
    return std::nullopt;
}

std::optional<User> User::getUserByUID(sql::Connection& con, int uid) {
    return std::nullopt;
}

// TODO: Validation methods
bool User::isEmailRegistered(sql::Connection& con, const std::string& email) {
    return false;
}

bool User::validatePassword(const std::string& password) const {
    return false;
}

std::string User::toString() const {
    std::ostringstream oss;
    oss << "User: {"
        << "uid: " << getUID() << ", "
        << "userName: " << getUserName() << ", "
        << "userPassword: " << getUserPassword() << ", "
        << "userSex: " << getUserSex() << ", "
        << "userAge: " << getUserAge() << ", "
        << "roleID: " << getRoleID() << ", "
        << "userEmail: " << getUserEmail() << ", "
        << "dateBirth: " << std::put_time(&dateBirth_, "%Y/%m/%d") << "}";
    return oss.str();
}