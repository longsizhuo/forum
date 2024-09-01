#include "user.hpp"
#include <bcrypt/BCrypt.hpp> // 使用 bcrypt-cpp 库进行密码哈希
#include <sstream>
#include <iomanip>
#include <string>
#include "service/database.hpp"

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
bool User::saveToDatabase(Database& db) const {
    return false;
}

std::optional<User> User::getUserByEmail(Database& db, const std::string& email) {
    return std::nullopt;
}

std::optional<User> User::getUserByUID(Database& db, int uid) {
    return std::nullopt;
}

// TODO: Validation methods
bool User::isEmailRegistered(Database& db, const std::string& email) {
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