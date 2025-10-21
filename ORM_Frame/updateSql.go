package main

func mainUpdate() {
	// -------------------------------------------------------------------------------Upsert 及 冲突 处理函数
	// GORM为不同数据库提供了对Upsert的兼容性支持
	// 存在冲突,创建终止
	//db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)

	// 在“id” 冲突时将列更新为默认值
	// db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "id"}},
	//	DoUpdates: clause.Assignments(map[string]interface{}{"role": "user"}),
	// }).Create(&users)

	// 使用合并为“用户”，如果未匹配，则插入，如果匹配，则更新设置
	// SQL如: INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `count`=GREATEST(count, VALUES(count)); ↓
	// db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "id"}},
	//	DoUpdates: clause.Assignments(map[string]interface{}{"count": gorm.Expr("GREATEST(count, VALUES(count))")}),
	// }).Create(&users)

	// 在“id”冲突时将列更新为新值
	// 使用合并为“用户”，如果未匹配，则插入，如果匹配，则更新设置“name”=“excluded”。“姓名”；SQL Server
	// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age"; PostgreSQL
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age`=VALUES(age);
	// db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "id"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	// }).Create(&users)

	// 当发生冲突时，将所有列更新为新值，但主键和sql func中具有默认值的列除外
	// SQL如: INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age", ...;
	// db.Clauses(clause.OnConflict{
	//	UpdateAll: true,
	// }).Create(&users)
}
