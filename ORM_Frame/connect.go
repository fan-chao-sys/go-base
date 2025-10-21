package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func Run(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Author{}, &Blog{}, &Blog2{}, &Blog3{})
	if err != nil {
		return
	}
}

func main() {
	// 连接数据库
	// 方法一：
	dbAddress := "root:root@tcp(127.0.0.1:3306)/go_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbAddress), &gorm.Config{})
	if err != nil {
	}
	Run(db)

	// 插入数据
	user := &User{}
	user.Age = 100
	user.MemberNumber.Valid = true // 字段值不明确情况下,指定设置需要写入空字符串,而非自动写入null
	db.Create(user)

	// 方法二：
	//_, _ = gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local",
	//	DefaultStringSize:         256,   // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})
}
