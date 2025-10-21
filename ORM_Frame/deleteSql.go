package main

import (
	"fmt"
	"gorm.io/gorm"
)

func mainDel() {
	// 删除一条记录
	// 删除一条记录时，删除对象需要指定主键，否则会触发 批量删除
	// Email 的 ID 是 `10`
	// DELETE from emails where id = 10;
	db.Delete(&email)
	// 带额外条件的删除
	// DELETE from emails where id = 10 AND name = "jinzhu";
	db.Where("name = ?", "jinzhu").Delete(&email)

	// 根据主键删除
	// 允许通过主键(可是复合主键)和内联条件来删除对象，它可以使用数字也可以使用字符串
	// DELETE FROM users WHERE id = 10;
	db.Delete(&User{}, 10)
	// DELETE FROM users WHERE id = 10;
	db.Delete(&User{}, "10")
	// DELETE FROM users WHERE id IN (1,2,3);
	db.Delete(&users, []int{1, 2, 3})

	// 批量删除
	// 如果指定的值不包括主属性，那么 GORM 会执行批量删除，它将删除所有匹配的记录
	// DELETE from emails where email LIKE "%jinzhu%";
	db.Where("email LIKE ?", "%jinzhu%").Delete(&Email{})
	// DELETE from emails where email LIKE "%jinzhu%";
	db.Delete(&Email{}, "email LIKE ?", "%jinzhu%")

	// 可以将一个主键切片传递给Delete 方法，以便更高效的删除数据量大的记录
	var users = []User{{ID: 1}, {ID: 2}, {ID: 3}}
	// DELETE FROM users WHERE id IN (1,2,3);
	db.Delete(&users)
	// DELETE FROM users WHERE name LIKE "%jinzhu%" AND id IN (1,2,3);
	db.Delete(&users, "name LIKE ?", "%jinzhu%")

	// 阻止全局删除
	// 当你试图执行不带任何条件的批量删除时，GORM将不会运行并返回ErrMissingWhereClause 错误
	// 如果一定要这么做，你必须添加一些条件，或者使用原生SQL，或者开启AllowGlobalUpdate 模式
	err1 := db.Delete(&User{}).Error                                       // gorm.ErrMissingWhereClause
	err2 := db.Delete(&[]User{{Name: "jinzhu1"}, {Name: "jinzhu2"}}).Error // gorm.ErrMissingWhereClause
	fmt.Println(err1, err2)

	// DELETE FROM `users` WHERE 1=1
	db.Where("1 = 1").Delete(&User{})
	// DELETE FROM users
	db.Exec("DELETE FROM users")
	// DELETE FROM users
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})

	// 软删除
	// 如果你的模型包含了 gorm.DeletedAt字段（该字段也被包含在gorm.Model中），那么该模型将会自动获得软删除的能力！
	// 当调用Delete时，GORM并不会从数据库中删除该记录，而是将该记录的DeleteAt设置为当前时间，而后的一般查询方法将无法查找到此条记录。
	// 查找被软删除的记录
	// 使用Unscoped来查询到被软删除的记录
	// SELECT * FROM users WHERE age = 20;
	db.Unscoped().Where("age = 20").Find(&users)

	// 永久删除
	// 使用 Unscoped来永久删除匹配的记录
	// DELETE FROM orders WHERE id=10;
	db.Unscoped().Delete(&order)

}
