// models/model.go
package models

import (
	"time"
)

// AuthRole 表示角色模型
type AuthRole struct {
	RoleID     int     `gorm:"primaryKey"`
	RoleName   string  `gorm:"size:20"`
	RoleStatus int     `gorm:"default:1"` // 0 = DELETE, 1 = ACTIVE, 2 = BLOCK
	Users      []User  `gorm:"foreignKey:RoleID"`
	Admins     []Admin `gorm:"foreignKey:RoleID"`
}

// User 表示用户模型
type User struct {
	UID            int       `gorm:"primaryKey"`
	UserBirth      string    `gorm:"size:20"`
	UserEmail      string    `gorm:"size:50"`
	UserName       string    `gorm:"size:20"`
	UserPassword   string    `gorm:"size:20"`
	UserSex        string    `gorm:"size:10"`
	UserAge        int       `gorm:"size:3"`
	UserOccupation string    `gorm:"size:50"`
	UserHobby      string    `gorm:"size:255"`
	DateBirth      time.Time `gorm:"type:date"`
	UserPoint      int       `gorm:"default:0"`
	UserClass      int       `gorm:"default:0"`
	UserRegister   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	RoleID         int       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleID"`
	Status         int       `gorm:"default:1"` // 0 = DELETE, 1 = ACTIVE, 2 = BLOCK
	Sections       []Section `gorm:"foreignKey:UID"`
	Topics         []Topic   `gorm:"foreignKey:UID"`
	Replies        []Reply   `gorm:"foreignKey:UID"`
}

// Admin 表示管理员模型
type Admin struct {
	AdminID       int    `gorm:"primaryKey"`
	AdminName     string `gorm:"size:20"`
	AdminPassword string `gorm:"size:20"`
	RoleID        int    `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleID"`
	Status        int    `gorm:"default:1"` // 0 = DELETE, 1 = ACTIVE, 2 = BLOCK
}

// Section 表示版块模型
type Section struct {
	SID        int    `gorm:"primaryKey"`
	UID        int    `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UID"` // Moderator (UserId)
	Name       string `gorm:"size:50"`
	Statement  string `gorm:"type:text"`
	ClickCount int
	TopicCount int
	Topics     []Topic `gorm:"foreignKey:SID"`
}

// Topic 表示主题模型
type Topic struct {
	TID        int    `gorm:"primaryKey"`
	SID        int    `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:SID"` // TopicSectionID
	UID        int    `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UID"` // TopicUserId
	Title      string `gorm:"size:20"`
	Content    string `gorm:"type:text"`
	Time       time.Time
	ReplyCount int
	ClickCount int
	Replies    []Reply `gorm:"foreignKey:TID"`
}

// Reply 表示回复模型
type Reply struct {
	RID        int    `gorm:"primaryKey"`
	TID        int    `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TID"` // ReplyTopicID
	UID        int    `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UID"` // ReplyUserId
	Content    string `gorm:"type:text"`
	Time       time.Time
	ClickCount int
}
