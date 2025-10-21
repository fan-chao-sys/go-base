package main

func mainCreate() {
	// 单项创建
	//result := db.Create(&user)                                                                           // 通过数据的指针来创建
	//fmt.Printf("单项创建结果,error:%s, rowsAffected：%d,返回主键:%d\n", result.Error, result.RowsAffected, user.ID) // ↓
	// user.ID                        // 返回插入数据的主键
	// result.Error                   // 返回 error
	// result.RowsAffected            // 返回插入记录的条数
	// user.MemberNumber.Valid = true // 字段值不明确情况下,指定设置需要写入空字符串,而非自动写入null

	// 指定字段赋值创建
	// 例如:INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
	//db.Select("Name", "Age", "Birthday").Create(&user)

	// 忽略传递给 ‘Omit’ 的字段值创建
	// 例如:INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
	//db.Omit("Name", "Age", "CreatedAt").Create(&user)

	// ****批量插入 : 切片传递给Create方法,GORM 将生成一条 SQL 来插入所有数据，以返回所有主键值，并触发 Hook 方法,将这些记录分割多批次,GORM一个事务去执行
	// 多项创建
	//db.Create(&users)
	//for _, user := range users {
	//	fmt.Println("users-打印创建后自增主键id:", user.ID)
	//}

	//db.Create(&userCopyArr)
	//for _, user: = range userCopyArr {
	//	fmt.Println("userCopyArr-打印创建后自增主键id:", user.ID)
	//}

	// 可通过 db.CreateInBatches方法来指定批量插入的批次大小

	//for i := 1; i <= 10; i++ {
	//	userBatch = append(userBatch, User{Name: fmt.Sprintf("jinzhu_%d", i)}) // 拼接值
	//}
	//// batch size 100
	//db.CreateInBatches(userBatch, 20)

	// ------------------------------------------------------------------------------- 创建钩子
	// SkipHooks会话模式, 跳过当前执行所有钩子方法
	// db.Session(&gorm.Session{SkipHooks: true}).Create(&user)
	// db.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(users, 100)

	// ------------------------------------------------------------------------------- 关联创建
	//db.Create(teacher)
	// 通过Select, Omit方法来跳过关联更新
	//db.Omit("CreditCard").Create(&user)
	// skip all associations
	//db.Omit (clause.Associations).Create(&user)
}
