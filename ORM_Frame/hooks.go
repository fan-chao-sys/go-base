package main

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

// *** 所有 Hooks 钩子函数的执行,都是针对某个固定结构体操作数据库时,自动触发的!

// BeforeSave 在创建或更新操作执行 SQL 语句之前触发。如果在这个钩子中返回错误，后续的创建或更新操作将被中断。例如，在保存用户信息前，可以利用这个钩子对用户输入的数据进行格式校验
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Age < 0 {
		return fmt.Errorf("age cannot be negative")
	}
	return nil
}

// BeforeCreate 在执行创建操作（INSERT 语句）之前触发。通常用于在创建对象时自动设置一些默认值，比如创建用户时自动生成注册时间。
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}

// AfterCreate 在创建操作（INSERT 语句）成功执行之后触发。可以用于记录日志、更新缓存等操作。比如在创建订单后，更新库存缓存。
func (o *Order) AfterCreate(tx *gorm.DB) (err error) {
	// 假设这里有一个更新库存缓存的函数
	//UpdateInventoryCache(o.ProductID, o.Quantity)
	return nil
}

// BeforeUpdate
// BeforeSave：如前所述，在更新操作执行 SQL 语句之前触发。可以用于验证更新数据是否符合要求，例如检查更新后的密码长度是否合规。
// BeforeUpdate：在执行更新操作（UPDATE 语句）之前触发，比BeforeSave更专注于更新操作。可以用于在更新前记录旧值，方便审计。
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	var oldUser User
	tx.First(&oldUser, u.ID)
	// 记录旧的用户信息到审计表
	//LogUserAudit(oldUser, *u)
	return nil
}

// AfterUpdate 在更新操作（UPDATE 语句）成功执行之后触发。可以用于通知相关系统数据已更新，例如在用户信息更新后通知消息推送系统。
func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	// 假设这里有一个通知消息推送系统的函数
	//NotifyMessageSystem(u.ID, "User information has been updated")
	return nil
}

// BeforeDelete 在执行删除操作（DELETE 语句）之前触发。可以用于实现软删除，即将数据标记为已删除而不是真正从数据库中删除。
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	u.DeletedAt = time.Now()
	return tx.Select("DeletedAt").Updates(u).Error
}

// AfterDelete 在删除操作（DELETE 语句）成功执行之后触发。可以用于清理相关的资源，比如删除用户后，删除该用户相关的文件存储记录。
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	// 假设这里有一个清理用户文件存储记录的函数
	//CleanUserFileRecords(u.ID)
	return nil
}

// AfterFind 在查询操作（SELECT 语句）成功执行之后触发。可以用于对查询结果进行进一步的处理，比如将查询到的时间字段进行格式化。
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.CreatedAt = u.CreatedAt.Local()
	return nil
}
