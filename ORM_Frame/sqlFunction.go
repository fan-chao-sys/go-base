package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var userArr []User
var user11 = User{ID: 1}
var name string
var age int
var total int64

// ------------------------------------------------------------------------------------------------------- sql函数使用介绍

func main() {
	// Raw():
	// 直接执行原始 SQL 语句，适合复杂查询场景
	// 执行原始查询并扫描结果到 users 切片
	db.Raw("SELECT * FROM users WHERE age > ?", 18).Scan(&userArr)

	// Exec() :
	// 执行写操作（INSERT/UPDATE/DELETE 等），返回 SQL 执行结果。
	result := db.Exec("UPDATE users SET age = ? WHERE id = ?", 20, 1)
	// 获取受影响的行数
	fmt.Println(result.RowsAffected)

	// Scan():
	// 将查询结果映射到结构体或变量中，常与 Raw() 配合使用
	// 扫描单条记录的特定字段
	db.Raw("SELECT name, age FROM users WHERE id = ?", 1).Scan(&user11)

	// Session():
	// 创建一个新的会话，可独立设置配置（如超时、是否跳过钩子等），不影响全局配置。
	// 创建一跳过钩子函数会话
	session := db.Session(&gorm.Session{SkipHooks: true})
	// 使用该会话执行操作（不会触发 BeforeSave/AfterSave 等钩子）
	session.Create(&User{Name: "test"})

	// Table():
	// 指定要操作的数据库表名（默认使用模型对应的表名，复数形式）
	// 显式指定表名为 "user_list"（而非默认的 "users"）
	db.Table("user_list").Where("status = ?", 1).Find(&users)

	// Select():
	// 指定查询返回的字段，或更新时的字段
	// 查询时只返回 name 和 age 字段
	db.Select("name", "age").Find(&users)
	// 更新时只更新 email 字段
	db.Model(&User{}).Where("id = ?", 1).Select("email").Updates(map[string]interface{}{
		"email": "new@example.com",
		"age":   30, // 此字段不会被更新，因为未在 Select 中指定
	})

	// Unscoped():
	// 取消软删除限制，查询 / 操作包含已软删除的记录（需模型启用软删除）
	// 默认查询（不包含软删除记录）
	db.Find(&users)
	// 包含软删除记录的查询
	db.Unscoped().Find(&users)
	// 永久删除记录（而非软删除）
	db.Unscoped().Delete(&User{}, 1)

	// Model():
	// 指定操作的模型（用于构建查询条件或关联操作），也可传入实例限定作用范围
	// 基于模型 User 构建查询
	db.Model(&User{}).Where("age > ?", 18).Find(&users)
	// 传入实例，限定操作当前实例（如更新该用户的信息）
	db.Model(&user11).Update("age", 25)

	// gorm.Expr():
	// 生成 SQL 表达式，用于复杂的计算或函数调用（如 COUNT()、NOW() 等）
	// 更新时使用表达式（自增 age 字段）
	db.Model(&User{}).Where("id = ?", 1).Update("age", gorm.Expr("age + ?", 1))
	// 查询时使用聚合函数
	db.Model(&User{}).Select("COUNT(*)").Scan(&total)

	// Clauses()：
	// 添加 SQL 子句（如 ON CONFLICT、FOR UPDATE 等），扩展查询功能
	// 插入时处理冲突（若 id 已存在则更新 name）
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name": "new_name"}),
	}).Create(&User{ID: 1, Name: "old_name"})

	// Take():
	// 查询单条记录（不保证顺序，等效于 LIMIT 1)
	// 查询任意一条满足条件的记录
	db.Where("status = ?", 1).Take(&user)

	// Find():
	// 查询多条记录，结果存入切片；若传入单个结构体，仅返回第一条
	// 查询所有年龄大于 18 的用户
	db.Where("age > ?", 18).Find(&users)

	// Rows():
	// 获取原生 SQL 结果集的行迭代器，适合处理大量数据或自定义扫描逻辑
	rows, err := db.Model(&User{}).Where("age > ?", 18).Rows()
	if err != nil {
	}
	defer rows.Close() // 必须调用！释放数据库连接
	// 迭代行数据
	for rows.Next() {
		// 扫描映射当前行到结构体
		db.ScanRows(rows, &user)
		// 处理 user
	}

	// Preload():
	// 程序执行前-预加载关联，一次性加载其关联的模型数据，避免多次数据库查询，提高效率
	db.Preload("user表外键关联字段名").First(&user, 1)
}
