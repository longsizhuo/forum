package models

import (
	"time"
)

// AuthRole 表示角色模型
type AuthRole struct {
	RoleID     int    `gorm:"primaryKey"`
	RoleName   string `gorm:"size:20"`
	RoleStatus int    `gorm:"default:1"` // 0 = DELETE, 1 = ACTIVE, 2 = BLOCK
}

// User 表示用户模型
type User struct {
	UID            int    `gorm:"primaryKey"`
	UserName       string `gorm:"size:20"`
	UserPassword   string `gorm:"size:20"`
	UserSex        string `gorm:"size:10"`
	UserAge        int
	UserOccupation string `gorm:"size:50"`
	UserHobby      string `gorm:"size:255"`
	UserPoint      int
	UserClass      int
	UserRegister   time.Time
	RoleID         int `gorm:"index"`
	Status         int `gorm:"default:1"` // 0 = DELETE, 1 = ACTIVE, 2 = BLOCK
}

// Admin 表示管理员模型
type Admin struct {
	AdminID       int    `gorm:"primaryKey"`
	AdminName     string `gorm:"size:20"`
	AdminPassword string `gorm:"size:20"`
	RoleID        int    `gorm:"index"`
	Status        int    `gorm:"default:1"` // 0 = DELETE, 1 = ACTIVE, 2 = BLOCK
}

// Section 表示版块模型
type Section struct {
	SID        int    `gorm:"primaryKey"`
	UID        int    `gorm:"index"` // Moderator (UserId)
	Name       string `gorm:"size:50"`
	Statement  string `gorm:"type:text"`
	ClickCount int
	TopicCount int
}

// Topic 表示主题模型
type Topic struct {
	TID        int    `gorm:"primaryKey"`
	SID        int    `gorm:"index"` // TopicSectionID
	UID        int    `gorm:"index"` // TopicUserId
	Title      string `gorm:"size:20"`
	Content    string `gorm:"type:text"`
	Time       time.Time
	ReplyCount int
	ClickCount int
}

// Reply 表示回复模型
type Reply struct {
	RID        int    `gorm:"primaryKey"`
	TID        int    `gorm:"index"` // ReplyTopicID
	UID        int    `gorm:"index"` // ReplyUserId
	Content    string `gorm:"type:text"`
	Time       time.Time
	ClickCount int
}
