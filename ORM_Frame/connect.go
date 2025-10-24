package main

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/hints"
	"net/http"
	"strconv"
	"time"
)

// 字段标签:
//	标签名									说明
//	column			指定 db 列名
//	type			列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用，例如：not null、size, autoIncrement… 像 varbinary(8) 这样指定数据库数据类型也是支持的。在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT
//	serializer		指定将数据序列化或反序列化到数据库中的序列化器, 例如: serializer:json/gob/unixtime
//	size			定义列数据类型的大小或长度，例如 size: 256
//	primaryKey		将列定义为主键
//	unique			将列定义为唯一键
//	default			定义列的默认值
//	precision		指定列的精度
//	scale			指定列大小
//	not null		指定列为 NOT NULL
//	autoIncrement	指定列为自动增长
//	autoIncrementIncrement	自动步长，控制连续记录之间的间隔
//	embedded		嵌套字段
//	embeddedPrefix	嵌入字段的列名前缀
//	autoCreateTime	创建时追踪当前时间，对于 int 字段，它会追踪时间戳秒数，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoCreateTime:nano
//	autoUpdateTime	创建/更新时追踪当前时间，对于 int 字段，它会追踪时间戳秒数，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoUpdateTime:milli
//	index			根据参数创建索引，多个字段使用相同的名称则创建复合索引，查看 索引 获取详情
//	uniqueIndex		与 index 相同，但创建的是唯一索引
//	check			创建检查约束，例如 check:age > 13，查看 约束 获取详情
//	<-				设置字段写入的权限， <-:create 只创建、<-:update 只更新、<-:false 无写入权限、<- 创建和更新权限
//	->				设置字段读的权限，->:false 无读权限
//	-				忽略该字段，- 表示无读写，-:migration 表示无迁移权限，-:all 表示无读写迁移权限
//	comment			迁移时为字段添加注释

// Run 驱动数据库方法
func Run(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Author{}, &Blog{}, &Blog2{}, &Blog3{})
	if err != nil {
		return
	}
}

func main321() {
	// --------------------------------------------------------------------------------------------    连接数据库
	// 方法一：自定义配置连接设置
	//_, _ = gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local",
	//	DefaultStringSize:         256,   // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})

	// 方法二：简便连接
	dbAddress := "root:root@tcp(127.0.0.1:3306)/go_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbAddress), &gorm.Config{})
	if err != nil {
	}
	Run(db)

	// -- 打印SQL语句日志 --
	db.Debug()

	// ----------------------------------------------------------------------------------------------- SQL构造器-原生写法
	// scan
	result := Result{}
	db.Raw("SELECT id, name, age FROM users WHERE id = ?", 3).Scan(&result)
	db.Raw("SELECT id, name, age FROM users WHERE name = ?", "jinzhu").Scan(&result)
	var age int
	db.Raw("SELECT SUM(age) FROM users WHERE role = ?", "admin").Scan(&age)
	var users []User
	db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "jinzhu", 20).Scan(&users)

	// Exec
	db.Exec("DROP TABLE users")
	db.Exec("UPDATE orders SET shipped_at = ? WHERE id IN ?", time.Now(), []int64{1, 2, 3})
	// Exec with SQL Expression
	db.Exec("UPDATE users SET money = ? WHERE name = ?", gorm.Expr("money * ? + ?", 10000, 1), "jinzhu")

	// DryRun 模式
	// 在不执行的情况下生成 SQL 及其参数，可以用于准备或测试生成的 SQL
	stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
	s := stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
	vars := stmt.Vars      //=> []interface{}{1}
	fmt.Println(s, vars)

	// Row & Rows
	// 获取 *sql.Row 结果
	// 用 GORM API 构建 SQL
	row := db.Table("users").Where("name = ?", "jinzhu").Select("name", "age").Row()
	name := ""
	err = row.Scan(&name, &age)
	if err != nil {
		return
	}
	// 用原生 SQL
	row = db.Raw("select name, age, email from users where name = ?", "jinzhu").Row()
	err = row.Scan(&name, &age, &email)
	if err != nil {
		return
	}

	// 获取 *sql.Rows 结果
	// 使用 GORM API 构建 SQL
	rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows()
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	for rows.Next() {
		err := rows.Scan(&name, &age, &email)
		if err != nil {
			return
		}
		// 业务逻辑...
	}

	// 原生 SQL
	rows, err = db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows()
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	for rows.Next() {
		err := rows.Scan(&name, &age, &email)
		if err != nil {
			return
		}
		// 业务逻辑...
	}

	// 将 sql.Rows 扫描至 model
	// 使用 ScanRows 将一行记录扫描至 struct
	rows, err = db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows() // (*sql.Rows, error)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	for rows.Next() {
		// ScanRows 将一行扫描至 user
		err := db.ScanRows(rows, &user)
		if err != nil {
			return
		}
		// 业务逻辑...
	}

	// ----------------------------------------------------------------------------------- 事务
	// 确保数据一致性，GORM在事务执行写入操作（创建、更新、删除）。可在初始化时禁用它，将获得大约 30%+ 性能提升。

	// 全局禁用
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{SkipDefaultTransaction: true})

	// 持续会话模式
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	tx.First(&user, 1)
	tx.Model(&user).Update("Age", 18)

	db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&User{Name: "Giraffe"}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		// 返回 nil 提交事务
		return nil
	})

	// 嵌套事务
	db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&user)
		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&user)
			return errors.New("rollback user") // Rollback user2
		})
		tx.Transaction(func(tx3 *gorm.DB) error {
			tx3.Create(&user)
			return nil
		})
		return nil
	})

	// 手动事务
	// 开始事务
	tx = db.Begin()
	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	tx.Create(&user)
	tx.Rollback()        // 回滚整个事务,所有操作全部撤回
	tx.SavePoint("sp1")  // 事务创建 'sp1' 保存点标记位置
	tx.RollbackTo("sp1") // 回退指定保存点位置 'sp1',只会撤销从该保存点之后 到 调 RollbackTo 之前的数据库操作，而保存点之前的操作仍然有效。
	// 否则，提交事务
	tx.Commit()

	// --------------------------------------------------------------------- 性能
	// GORM 已优化许多提高性能，另一些关于应用改进性能的方法:
	// *禁用默认事务
	// *缓存预编译语句
	// 全局模式
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		PrepareStmt: true,
	})
	// 会话模式
	tx = db.Session(&gorm.Session{PrepareStmt: true})
	tx.First(&user, 1)
	tx.Model(&user).Update("Age", 18)
	// *选择字段
	db.Select("Name", "Age").Find(&User{})
	// *自主选择索引
	db.Clauses(hints.UseIndex("idx_user_name")).Find(&User{})

	// ------------------------------------------------------------------------Scopes
	// 作用域允许你复用通用的逻辑，这种共享逻辑需要定义为类型func(*gorm.DB) *gorm.DB
	db.Scopes(AmountGreaterThan1000, PaidWithCreditCard).Find(&order)

	// 分页
	r, _ := http.NewRequest("GET", "http://example.com/some/path?page=2&page_size=20", nil)
	db.Scopes(Paginate(r)).Find(&user)
	db.Scopes(Paginate(r)).Find(&order)

	// 指定表
	// Table form different database
	// SELECT * FROM users_2019;
	db.Scopes(TableOfYear(&user, 2019)).Find(&user)
	// SELECT * FROM users_2020;
	db.Scopes(TableOfYear(&user, 2020)).Find(&user)
	// SELECT * FROM org1.users;
	db.Scopes(TableOfOrg(&user, "org1")).Find(&user)
	// SELECT * FROM org2.users;
	db.Scopes(TableOfOrg(&user, "org2")).Find(&user)

	// 动态条件
	// UPDATE articles SET name = "name 1" WHERE org_id = 111
	db.Model(&order).Scopes(CurOrganization(r)).Update("Name", "name 1")
	// DELETE FROM articles WHERE org_id = 111
	db.Scopes(CurOrganization(r)).Delete(&order)

	// ----------------------------------------------------------------------------------------- 设置
	// Set, Get, InstanceSet, InstanceGet 方法来允许用户传值给 勾子 或其他方法
	myValue := 123
	db.Set("my_value", myValue).Create(&User{})
	// // 在创建关联时，GORM 创建了一个新 `*Statement`，所以它不能读取到其它实例的设置
	myValueTwo := 123
	db.InstanceSet("my_value", myValueTwo).Create(&User{})
}

// AmountGreaterThan1000 作用域 -- scopes
func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
	return db.Where("amount > ?", 1000)
}

// PaidWithCreditCard 作用域 -- scopes
func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
	return db.Where("pay_mode_sign = ?", "C")
}

// PaidWithCod 作用域 -- scopes
func PaidWithCod(db *gorm.DB) *gorm.DB {
	return db.Where("pay_mode_sign = ?", "C")
}

// OrderStatus 作用域 -- scopes
func OrderStatus(status []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status IN (?)", status)
	}
}

// Paginate 分页
func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// TableOfYear 指定表
func TableOfYear(user *User, year int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := user.Name + strconv.Itoa(year)
		return db.Table(tableName)
	}
}

// TableOfOrg 指定表
func TableOfOrg(user *User, dbName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := dbName + "." + user.Name
		return db.Table(tableName)
	}
}

// CurOrganization 动态条件
func CurOrganization(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		org := r.Query("org")

		if org != "" {
			var teacher Teacher
			if db.Session(&gorm.Session{}).First(&teacher, "name = ?", org).Error == nil {
				return db.Where("org_id = ?", teacher.ID)
			}
		}
		return db
	}
}

type UserT struct {
	gorm.Model
	CreditCard CreditCard
	// ...
}

func (u *UserT) BeforeCreate(tx *gorm.DB) error {
	myValue, ok := tx.Get("my_value")
	// ok => true
	if ok {
	}
	// myValue => 123
	if myValue == 123 {
	}
	return nil
}

type CreditCardT struct {
	gorm.Model
	// ...
}

func (card *CreditCardT) BeforeCreate(tx *gorm.DB) error {
	myValue, ok := tx.Get("my_value")
	// ok => true
	if ok {
	}
	// myValue => 123
	if myValue == 123 {
	}
	return nil
}

// BeforeCreate 在创建关联时，GORM 创建了一个新 `*Statement`，所以它不能读取到其它实例的设置
func (card *CreditCard) BeforeCreate(tx *gorm.DB) error {
	myValue, ok := tx.InstanceGet("my_value")
	// ok => false
	if ok {
	}
	// myValue => 123
	if myValue == 123 {
	}
	return nil
}
