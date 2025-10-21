package main

func mainQuery() {
	// --------------------------------------------------------------------------------- 查询
	// 检索单个对象
	// GORM 提供了 First、Take、Last 方法，以便从数据库中检索单个对象。当查询数据库时它添加了 LIMIT 1 条件，且没有找到记录时，它会返回 ErrRecordNotFound 错误
	// 获取第一条记录（主键升序）
	// SELECT * FROM users ORDER BY id LIMIT 1;
	// db.First(&user)

	// 获取一条记录，没有指定排序字段
	// SELECT * FROM users LIMIT 1;
	// db.Take(&user)

	// 获取最后一条记录（主键降序）
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	// db.Last(&user)

	// 检查 ErrRecordNotFound 错误
	// result := db.First(&user)
	// errors.Is(result.Error, gorm.ErrRecordNotFound)

	// 根据主键检索
	// 如果主键是数字类型，可用 内联条件 检索对象。 当使用字符串时，需要额外的注意来避免SQL注入；
	// SELECT * FROM users WHERE id = 10;
	//db.First(&user, 10)

	// SELECT * FROM users WHERE id = 10;
	//db.First(&user, "10")

	// SELECT * FROM users WHERE id IN (1,2,3);
	//db.Find(&users, []int{1, 2, 3})

	// 如果主键是字符串(例如像uuid)，查询将被写成如下
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";
	//db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")

	// 当目标对象有一个主键值时，将使用主键构建查询条件
	// SELECT * FROM users WHERE id = 10;
	// var u = User{ID: 10}
	// db.First(&u)

	// SELECT * FROM users WHERE id = 10;
	// var result User
	// db.Model(User{ID: 10}).First(&result)

	// 检索全部对象
	// SELECT * FROM users;
	//db.Find(&user)

	// String 条件
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	//db.Where("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE name <> 'jinzhu';
	//db.Where("name <> ?", "jinzhu").Find(&user)
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');
	//db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&user)
	// SELECT * FROM users WHERE name LIKE '%jin%';
	//db.Where("name LIKE ?", "%jin%").Find(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&user)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
	//db.Where("updated_at > ?", lastWeek).Find(&user)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
	//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&user)

	// 如果对象设置了主键，条件查询将不会覆盖主键的值，而是用 And 连接条件
	// 这个查询将会给出record not found错误 所以，在你想要使用例如 user 这样的变量从数据库中获取新值前，需要将例如 id 这样的主键设置为nil。
	//var user = User{ID: 10}
	//user.ID = nil
	// SELECT * FROM users WHERE id = 10 and id = 20 ORDER BY id ASC LIMIT 1
	//db.Where("id = ?", 20).First(&user)

	// Struct & Map 条件
	// Struct
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
	//db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	// Map
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);
	//db.Where([]int64{20, 21, 22}).Find(&users)

	// 指定结构体查询字段
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
	//db.Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)
	// SELECT * FROM users WHERE age = 0;
	//db.Where(&User{Name: "jinzhu"}, "Age").Find(&users)

	// 内联条件
	// 如果主键是非整数类型，则按主键获取
	// SELECT * FROM users WHERE id = 'string_primary_key';
	//db.First(&user, "id = ?", "string_primary_key")
	// SELECT * FROM users WHERE name = "jinzhu";
	//db.Find(&user, "name = ?", "jinzhu")
	// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;
	//db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT * FROM users WHERE age = 20;
	//db.Find(&users, User{Age: 20})
	// Map
	// SELECT * FROM users WHERE age = 20;
	//db.Find(&users, map[string]interface{}{"age": 20})

	// Not 条件
	// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;
	//db.Not("name = ?", "jinzhu").First(&user)
	// Not In
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");
	//db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;
	//db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// Not In slice of primary keys
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
	//db.Not([]int64{1, 2, 3}).First(&user)

	// Or 条件
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';
	//db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// Struct
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);
	//db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
	// Map
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);
	//db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)

	// 排序
	// SELECT * FROM users ORDER BY age desc, name;
	//db.Order("age desc, name").Find(&users)
	// Multiple orders
	// SELECT * FROM users ORDER BY age desc, name;
	//db.Order("age desc").Order("name").Find(&users)
	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
	//db.Clauses(clause.OrderBy{
	//	Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	//}).Find(&User{})

	// Limit & Offset
	// SELECT * FROM users LIMIT 3;
	//db.Limit(3).Find(&users)
	// Cancel limit condition with -1
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)
	//db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users OFFSET 3;
	//db.Offset(3).Find(&users)
	// SELECT * FROM users OFFSET 5 LIMIT 10;
	//db.Limit(10).Offset(5).Find(&users)
	// Cancel offset condition with -1
	// SELECT * FROM users; (users2)
	// SELECT * FROM users OFFSET 10; (users1)
	//db.Offset(10).Find(&users1).Offset(-1).Find(&users2)

	// Group By & Having
	// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1
	//db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
	// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"
	//db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
	// 示例:
	//rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
	//defer rows.Close()
	//for rows.Next() {
	//}
	// db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)

	// Distinct
	//db.Distinct("name", "age").Order("name, age desc").Find(&results)

	// Joins
	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id
	//db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
	//rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
	//for rows.Next() {
	//}
	//db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	// multiple joins with parameter
	//db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)

	// Joins 预加载
	// 可以将Joins与单个SQL结合使用
	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;
	//db.Joins("Company").Find(&users)
	// inner join
	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` INNER JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;
	//db.InnerJoins("Company").Find(&users)

	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id` AND `Company`.`alive` = true;
	//db.Joins("Company", db.Where(&Company{Alive: true})).Find(&users)

	// Joins 一个衍生表
	// 可用联接来联接派生表。
	// SELECT `order`.`user_id`,`order`.`finished_at` FROM `order` join (SELECT MAX(order.finished_at) as latest FROM `order` left join user user on order.user_id = user.id WHERE user.age > 18 GROUP BY `order`.`user_id`) q on order.finished_at = q.latest
	//query := db.Table("order").Select("MAX(order.finished_at) as latest").Joins("left join user user on order.user_id = user.id").Where("user.age > ?", 18).Group("order.user_id")
	//db.Model(&Order{}).Joins("join (?) q on order.finished_at = q.latest", query).Scan(&results)

	// Scan： 将结果扫描到结构体中, 与使用Find方式类似
	//var result Result
	//db.Table("users").Select("name", "age").Where("name = ?", "Antonio").Scan(&result)
	// Raw SQL
	//db.Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)

}
