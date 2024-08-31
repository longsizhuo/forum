#include <string>
#include <ctime>

namespace sqlDb
{
    class AuthRole {
     public:
        int RoleID;
        std::string RoleName;
        int RoleStatus; // 0 = DELETE, 1 = ACTIVE, 2 = BLOCK
        std::string createAuthRoleTable = R"(
CREATE TABLE IF NOT EXISTS AuthRole (
    RoleID INT PRIMARY KEY AUTO_INCREMENT,
    RoleName VARCHAR(20),
    RoleStatus INT DEFAULT 1
);
        )";
        // Users 和 Admins 需要单独处理关联关系
    };

    class User {
     public:
        int UID;
        std::string UserBirth;
        std::string UserEmail;
        std::string UserName;
        std::string UserPassword;
        std::string UserSex;
        int UserAge;
        std::string UserOccupation;
        std::string UserHobby;
        std::time_t DateBirth;
        int UserPoint;
        int UserClass;
        int RoleID;
        int Status; // 0 = DELETE, 1 = ACTIVE, 2 = BLOCK
        // Sections, Topics, Replies 需要单独处理关联关系
        std::string createUserTable = R"(
CREATE TABLE IF NOT EXISTS User (
    UID INT PRIMARY KEY AUTO_INCREMENT,
    UserBirth VARCHAR(20),
    UserEmail VARCHAR(50),
    UserName VARCHAR(20),
    UserPassword VARCHAR(64),
    UserSex VARCHAR(10),
    UserAge INT,
    UserOccupation VARCHAR(50),
    UserHobby VARCHAR(255),
    DateBirth DATE,
    UserPoint INT DEFAULT 0,
    UserClass INT DEFAULT 0,
    RoleID INT NULL,
    Status INT DEFAULT 1,
    FOREIGN KEY (RoleID) REFERENCES AuthRole(RoleID) ON UPDATE CASCADE ON DELETE SET NULL
);
        )";
    };

    class Admin {
     public:
        int AdminID;
        std::string AdminName;
        std::string AdminPassword;
        int RoleID;
        int Status; // 0 = DELETE, 1 = ACTIVE, 2 = BLOCK
        std::string createAdminTable = R"(
CREATE TABLE IF NOT EXISTS Admin (
    AdminID INT PRIMARY KEY AUTO_INCREMENT,
    AdminName VARCHAR(20),
    AdminPassword VARCHAR(64),
    RoleID INT NULL,
    Status INT DEFAULT 1,
    FOREIGN KEY (RoleID) REFERENCES AuthRole(RoleID) ON UPDATE CASCADE ON DELETE SET NULL
);
        )";
    };

    class Section {
     public:
        int SID;
        int UID; // Moderator (UserId)
        std::string Name;
        std::string Statement;
        int ClickCount;
        int TopicCount;
        std::string createSectionTable = R"(
CREATE TABLE IF NOT EXISTS Section (
    SID INT PRIMARY KEY AUTO_INCREMENT,
    UID INT NULL,
    Name VARCHAR(50),
    Statement TEXT,
    ClickCount INT DEFAULT 0,
    TopicCount INT DEFAULT 0,
    FOREIGN KEY (UID) REFERENCES User(UID) ON UPDATE CASCADE ON DELETE SET NULL
);
        )";
        // Topics 需要单独处理关联关系
    };

    class Topic {
     public:
        int TID;
        int SID; // TopicSectionID
        int UID; // TopicUserId
        std::string Title;
        std::string Content;
        std::time_t Time;
        int ReplyCount;
        int ClickCount;
        // Replies 需要单独处理关联关系
        std::string createTopicTable = R"(
CREATE TABLE IF NOT EXISTS Topic (
    TID INT PRIMARY KEY AUTO_INCREMENT,
    SID INT NULL,
    UID INT NULL,
    Title VARCHAR(50),
    Content TEXT,
    Time DATETIME,
    ReplyCount INT DEFAULT 0,
    ClickCount INT DEFAULT 0,
    FOREIGN KEY (SID) REFERENCES Section(SID) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (UID) REFERENCES User(UID) ON UPDATE CASCADE ON DELETE SET NULL
);
        )";
    };

    class Reply {
     public:
        int RID;
        int TID; // ReplyTopicID
        int UID; // ReplyUserId
        std::string Content;
        std::time_t Time;
        int ClickCount;
        std::string createReplyTable = R"(
CREATE TABLE IF NOT EXISTS Reply (
    RID INT PRIMARY KEY AUTO_INCREMENT,
    TID INT NULL,
    UID INT NULL,
    Content TEXT,
    Time DATETIME,
    ClickCount INT DEFAULT 0,
    FOREIGN KEY (TID) REFERENCES Topic(TID) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (UID) REFERENCES User(UID) ON UPDATE CASCADE ON DELETE SET NULL
);
        )";
    };

    class Channel {
     public:
        int CID;
        // Users 需要单独处理关联关系
        std::string createChannelTable = R"(
CREATE TABLE IF NOT EXISTS Channel (
    CID INT PRIMARY KEY AUTO_INCREMENT,
    UID INT NULL,
    FOREIGN KEY (UID) REFERENCES User(UID) ON UPDATE CASCADE ON DELETE SET NULL
);
        )";
    };

    class Message {
     public:
        int MID;
        int CID; // 频道 ID
        int UID; // 用户 ID
        std::string Content; // 文本消息内容
        int* MediaID; // 如果是媒体消息，关联到媒体表
        std::time_t CreatedAt; // 消息发送时间
        std::string createMessageTable = R"(
CREATE TABLE IF NOT EXISTS Message (
    MID INT PRIMARY KEY AUTO_INCREMENT,
    CID INT NULL,
    UID INT NULL,
    Content TEXT,
    MediaID INT NULL,
    CreatedAt DATETIME
);
        )";
    };

    class Media {
     public:
        int MediaID;
        std::string FileType;  // 文件类型，例如：image, video, file
        std::string FilePath; // 文件路径或 URL
        int FileSize;   // 文件大小
        int UploadedBy; // 上传的用户 ID
        std::time_t UploadedAt; // 上传时间
        std::string createMediaTable = R"(
CREATE TABLE IF NOT EXISTS Media (
    MediaID INT PRIMARY KEY AUTO_INCREMENT,
    FileType VARCHAR(20),
    FilePath VARCHAR(255),
    FileSize INT,
    UploadedBy INT NULL,
    UploadedAt DATETIME,
    FOREIGN KEY (UploadedBy) REFERENCES User(UID) ON UPDATE CASCADE ON DELETE SET NULL
);
        )";
    };
}