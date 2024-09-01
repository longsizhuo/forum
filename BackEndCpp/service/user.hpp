#include <string>
#include <ctime>
#include <optional>
#include <iostream>

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
    bool saveToDatabase(Database& db) const;
    static std::optional<User> getUserByEmail(Database& db, const std::string& email);
    static std::optional<User> getUserByUID(Database& db, int uid);

    // 验证方法
    static bool isEmailRegistered(Database& db, const std::string& email);
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