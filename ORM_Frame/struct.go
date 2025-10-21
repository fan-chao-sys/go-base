package main

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// User 模型定义
type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

// Model 的定义，GORM提供了一个预定义的结构体，名为gorm.Model，其中包含常用字段：
// 将其嵌入在您的结构体中: 您可以直接在您的结构体中嵌入 gorm.Model ，以便自动包含这些字段
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Author 嵌入结构体,对于匿名字段，GORM 会将其字段包含在父结构体中
type Author struct {
	// Model              // 用户自定义的结构体嵌入写法方式
	gorm.Model // GORM 库自带的一个基础模型结构体
	Name       string
	Email      string
}

type Blog struct {
	Author
	ID      int
	Upvotes int32
}

// Author2 正常的结构体字段，你也可以通过标签 embedded 将其嵌入
type Author2 struct {
	Name  string
	Email string
}

type Blog2 struct {
	ID      int
	Author  Author `gorm:"embedded"` // gorm:字段标签配置！ embedded：当前对象中所有字段以嵌入方式全部创建出字段
	Upvotes int32  `gorm:"column:votes"`
}

// Blog3 可以使用标签 embeddedPrefix 来为 db 中的字段名添加前缀
type Blog3 struct {
	ID      int
	Author  Author `gorm:"embedded;embeddedPrefix:author_"`
	Upvotes int32
}
