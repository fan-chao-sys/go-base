package main

import (
	"fmt"
	"gorm.io/gorm"
)

func mainSave() {
	// save 是一个组合函数。 如果保存值不包含主键，它将执行 Create，否则它将执行 Update (包含所有字段)。
	// UPDATE users SET the name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
	db.First(&user)
	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)

	// 更新单个列
	// 当使用 Update 更新单列时，需要有一些条件，否则将会引起ErrMissingWhereClause 错误，当使用 Model 方法，并且它有主键值时，主键将会被用于构建条件
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;
	db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	// User 的 ID 是 `111`
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Update("name", "hello")
	// 根据条件和 model 的值进行更新
	// UPDATE users SET the name='hello ', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
	db.Model(&user).Where("active = ?", true).Update("name", "hello")

	// 更新多列
	// Updates 方法支持 struct 和 map[string]interface{} 参数。当使用 struct 更新时，默认情况下GORM 只会更新非零值的字段
	// 根据 `struct` 更新属性，只会更新非零值的字段
	// UPDATE users SET the name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
	db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	// 根据 `map` 更新属性
	// UPDATE users SET the name='hello ', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})

	// 更新选定字段
	// 如果想在更新选择、忽略某些字段，可用 Select、Omit
	// 选择 Map 的字段
	// User 的 ID 是 `111`:
	// UPDATE users SET name='hello' WHERE id=111;
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// 选择 Struct 的字段（会选中零值的字段）
	// UPDATE users' SET name='new_name', age=0 WHERE id=111;
	db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
	// 选择所有字段（选择包括零值字段的所有字段）
	db.Model(&user).Select("*").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})
	// 选择除 Role 外的所有字段（包括零值字段的所有字段）
	db.Model(&user).Select("*").Omit("Role").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})

	// 批量更新
	// 如果没有通过 Model 指定一个含有主键的记录，GORM 会执行批量更新
	// struct
	// UPDATE users SET name='hello', age=18 WHERE role = 'admin';
	db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
	// map
	// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
	db.Table("users").Where("id IN ?", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})

	// 阻止全局更新
	// 如果你执行一个没有任何条件的批量更新，GORM 默认不会运行，并且会返回 ErrMissingWhereClause 错误
	// 可用一些条件，用原生 SQL 或者启用 AllowGlobalUpdate 模式，例如：// gorm.ErrMissingWhereClause
	err := db.Model(&User{}).Update("name", "jinzhu").Error
	fmt.Println(err)
	// UPDATE users SET `name` = "jinzhu" WHERE 1=1
	db.Model(&User{}).Where("1 = 1").Update("name", "jinzhu")
	// UPDATE users SET name = "jinzhu"
	db.Exec("UPDATE users SET name = ?", "jinzhu")
	// UPDATE users SET `name` = "jinzhu"
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("name", "jinzhu")

	// 使用 SQL 表达式更新
	// UPDATE "User" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;
	db.Model(&User{}).Update("price", gorm.Expr("price * ? + ?", 2, 100))
	// UPDATE "User" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;
	db.Model(&User{}).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
	// UPDATE "User" SET "quantity" = quantity - 1 WHERE "id" = 3;
	db.Model(&User{}).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	// UPDATE "User" SET "quantity" = quantity - 1 WHERE "id" = 3 AND quantity > 1;
	db.Model(&User{}).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))

	// 根据子查询进行更新
	// UPDATE "users" SET "company_name" = (SELECT name FROM companies WHERE companies.id = users.company_id);
	db.Model(&user).Update("company_name", db.Model(&Company{}).Select("name").Where("companies.id = users.company_id"))
	db.Table("users as u").Where("name = ?", "jinzhu").Update("company_name", db.Table("companies as c").Select("name").Where("c.id = u.company_id"))
	db.Table("users as u").Where("name = ?", "jinzhu").Updates(map[string]interface{}{"company_name": db.Table("companies as c").Select("name").Where("c.id = u.company_id")})

	// 不使用 Hook 和时间追踪
	// 希望更新时跳过 Hook 方法，并不追踪更新时间，可用 UpdateColumn, UpdateColumns, 用法类似于 Update, Updates
	// Update single column
	// UPDATE users SET name='hello' WHERE id = 111;
	db.Model(&user).UpdateColumn("name", "hello")
	// Update multiple columns
	// UPDATE users SET name='hello', age=18 WHERE id = 111;
	db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
	// Update selected columns
	// UPDATE users SET name='hello', age=0 WHERE id = 111;
	db.Model(&user).Select("name", "age").UpdateColumns(User{Name: "hello", Age: 0})

}
