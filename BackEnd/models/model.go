// models/model.go
package models

import (
	"gorm.io/gorm"
	"time"
)

// User 表示用户模型
type User struct {
	gorm.Model
	UserBirth      string    `gorm:"size:20"`
	UserEmail      string    `gorm:"size:50"`
	UserName       string    `gorm:"size:20"`
	UserPassword   string    `gorm:"size:200"`
	UserSex        string    `gorm:"size:10"`
	UserAge        int       `gorm:"size:3"`
	UserOccupation string    `gorm:"size:50"`
	UserHobby      string    `gorm:"size:255"`
	DateBirth      time.Time `gorm:"type:date"`
	UserPoint      int       `gorm:"default:0"`
	UserClass      int       `gorm:"default:0"`
	RoleID         int       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleID"`
	Status         int       `gorm:"default:1"` // 0 = DELETE, 1 = ACTIVE, 2 = BLOCK
	Sections       []Section `gorm:"foreignKey:ID"`
	Topics         []Topic   `gorm:"foreignKey:ID"`
	Replies        []Reply   `gorm:"foreignKey:ID"`
	Friends        []User    `gorm:"many2many:user_friends"`
}

// Admin 表示管理员模型
type Admin struct {
	gorm.Model
	AdminName     string `gorm:"size:20"`
	AdminPassword string `gorm:"size:20"`
	RoleID        int    `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleID"`
	Status        int    `gorm:"default:1"` // 0 = DELETE, 1 = ACTIVE, 2 = BLOCK
}

// Section 表示版块模型
type Section struct {
	gorm.Model
	UID        int    `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UID"` // Moderator (UserId)
	Name       string `gorm:"size:50"`
	Statement  string `gorm:"type:text"`
	ClickCount int
	TopicCount int
	Topics     []Topic `gorm:"foreignKey:SID"`
}

// Topic 表示主题模型
type Topic struct {
	gorm.Model
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
	gorm.Model
	TID        int    `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TID"` // ReplyTopicID
	UID        int    `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UID"` // ReplyUserId
	Content    string `gorm:"type:text"`
	Time       time.Time
	ClickCount int
}

// Channel 表示频道模型
type Channel struct {
	gorm.Model
	Users []User `gorm:"foreignKey:CID"`
}

// Message 表示消息模型
type Message struct {
	gorm.Model
	CID       int       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CID"`     // 频道 ID
	UID       int       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UID"`     // 用户 ID
	Content   string    `gorm:"type:text"`                                                              // 文本消息内容
	MediaID   *int      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:MediaID"` // 如果是媒体消息，关联到媒体表
	CreatedAt time.Time // 消息发送时间
}

// Media 表示媒体模型
type Media struct {
	gorm.Model
	FileType   string    `gorm:"size:20"`  // 文件类型，例如：image, video, file
	FilePath   string    `gorm:"size:255"` // 文件路径或 URL
	FileSize   int       // 文件大小
	UploadedBy int       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UID"` // 上传的用户 ID
	UploadedAt time.Time // 上传时间
}
