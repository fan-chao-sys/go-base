package main

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

// foreignKey：指定关联模型（被依赖的 “子模型”） 中存储关联标识的字段（即外键字段本身）。
// references：指定当前模型（主动关联的 “主模型”） 中被关联的字段（即外键要指向的目标字段）。
// 用 “父子模型” 类比理解:
// 假设我们有两个模型 A 和 B，A 关联 B（A 是 “主”，B 是 “子”）：
//  foreignKey 是 “子模型 B” 身上的 “身份证”，用来记录 “主模型 A” 的标识。
// 	references 是 “主模型 A” 身上被记录的 “标识字段”（通常是主键或唯一字段）。
// polymorphic: 开启多态关联
// Owner： 多态关联标识
// 固定定义写法:
//  *polymorphic: 同时指定type和ID（value+Type value+ID）
//  *polymorphicType: 单独指定type
//  *polymorphicId: 单独指定ID
//  *polymorphicValue: 指定type的值，默认为表名

type Dog struct {
	ID   int
	Name string
	Toys []Toy `gorm:"polymorphic:Owner;"`
}

type Cat struct {
	ID   int
	Name string
	Toys []Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

// Session 配置: 允许创建带配置的新建会话模式
type Session struct {
	DryRun                   bool
	PrepareStmt              bool
	NewDB                    bool
	Initialized              bool
	SkipHooks                bool
	SkipDefaultTransaction   bool
	DisableNestedTransaction bool
	AllowGlobalUpdate        bool
	FullSaveAssociations     bool
	QueryFields              bool
	Context                  context.Context
	Logger                   logger.Interface
	NowFunc                  func() time.Time
	CreateBatchSize          int
}

// 自定义类型
type strArr []string

// UserD 自定义数据类型
type UserD struct {
	ID     int
	Skills strArr
}

func (arr strArr) Value() (driver.Value, error) {
	return strings.Join(arr, ","), nil
}

func (arr *strArr) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	*arr = strings.Split(string(bytes), ",")
	return nil
}

// Post 有许多第三方包实现了 Scanner/Valuer 接口，可与 GORM 一起使用，例如：
type Post struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title string
	Tags  pq.StringArray `gorm:"type:text[]"`
}

// UserOne  外键关联结构体关系
// FOREIGN KEY: 一个表中的 FOREIGN KEY 指向另一个表中的 UNIQUE KEY(唯一约束的键)。
// REFERENCES: 确定外键指向的另外一个表的字段
// belongs to 会与另一个模型建立了一对一的连接。 这种模型的每一个实例都“属于”另一个模型的一个实例。
// 例如，user 和 company，并且每个 user 能且只能被分配给一个 company。下面的类型就表示这种关系。
// 注意，User对象中，有一个和 Company 一样的 CompanyID。 默认情况下， CompanyID 被隐含地用来在 User 和 Company 之间创建一个外键关系， 因此必须包含在 User 结构体中才能填充 Company 内部结构体。
// `UserOne` 属于 `Company`，`CompanyID` 是外键
type UserOne struct {
	gorm.Model
	Name      string
	CompanyID int // 默认外键
	Company   Company
}

type Company struct {
	ID    int
	Name  string
	Alive bool
}

// UserTwo 重写外键
// 定义一个 belongs to 关系，数据库表中必须存在外键。默认情况下，外键的名字，使用拥有者的类型名称加上表的主键的字段名字
// 例如，定义一个User实体属于Company实体，那么外键的名字一般使用CompanyID。
// GORM同时提供自定义外键名字的方式，如下例所示。
type UserTwo struct {
	gorm.Model
	Name         string
	CompanyRefer int
	Company      Company `gorm:"foreignKey:CompanyRefer"` // 使用 CompanyRefer 作为外键
}

// User3 重写引用
// 主表（拥有者）主键值作外键参考。 正如上面例子，使用主表Company中的主键字段ID作为外键的参考值。
// 如果设置了User实体属于Company实体，那么GORM会自动把Company中的ID属性保存到User的CompanyID属性中。
// 同样的，可用标签 references 来更改它，例如：
type User3 struct {
	gorm.Model
	Name      string
	CompanyID string
	Company1  Company1 `gorm:"references:Code"` // 使用 Code 作为引用
}

type Company1 struct {
	ID   int
	Code string
	Name string
}

// User4 如果外键名恰好在拥有者类型中存在，GORM 通常会错误的认为它是 has one 关系。我们需要在 belongs to 关系中指定 references
type User4 struct {
	gorm.Model
	Name      string
	CompanyID string
	Company2  Company2 `gorm:"references:CompanyID"` // 使用 Company.CompanyID 作为引用
}

type Company2 struct {
	CompanyID int
	Code      string
	Name      string
}

// User5 ------------------------------------------------------------  一对一
// has one 与另一个模型建立一对一的关联，但它和一对一关系有些许不同。 这种关联表明一个模型的每个实例都包含或拥有另一个模型的一个实例。
// User 有一张 CreditCard，UserID 是外键
type User5 struct {
	gorm.Model
	CreditCard1 CreditCard1
}

type CreditCard1 struct {
	gorm.Model
	UserName string
	Number   string
	UserID   uint
}

// GetAll 检索用户列表并预加载信用卡
func GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("CreditCard1").Find(&users).Error
	return users, err
}

// User6 重写外键
// 对于 has one 关系，同样必须存在外键字段。拥有者将把属于它的模型的主键保存到这个字段。
// 这个字段的名称通常由 has one 模型的类型加上其 主键 生成，对于上面的例子，它是 UserID。
// 为 user 添加 credit card 时，它会将 user 的 ID 保存到自己的 UserID 字段。
// 如果你想要使用另一个字段来保存该关系，你同样可以使用标签 foreignKey 来更改它，例如：
type User6 struct {
	gorm.Model
	CreditCard1 CreditCard1 `gorm:"foreignKey:UserName"` // 使用 UserName 作为外键
}

// User7 重写引用
// 默认情况下，拥有者实体会将 has one 对应模型的主键保存为外键，您也可以修改它，用另一个字段来保存，例如下面这个使用 Name 来保存的例子。
// 使用标签 references 来更改它
type User7 struct {
	gorm.Model
	Name        string      `gorm:"index"`
	CreditCard1 CreditCard1 `gorm:"foreignKey:UserName;references:Name"`
}

// User8 一对多
// has many 与另一个模型建立了一对多的连接。 不同于 has one，拥有者可以有零或多个关联模型。
// 例如，您的应用包含 user 和 credit card 模型，且每个 user 可以有多张 credit card。
// User 有多张 CreditCard，UserID 是外键

// User8 重写外键
// 要定义 has many 关系，同样必须存在外键。 默认的外键名是拥有者的类型名加上其主键字段名
// 例如，要定义一个属于 User 的模型，则其外键应该是 UserID。
// 此外，想要使用另一个字段作为外键，您可以使用 foreignKey 标签自定义它：
type User8 struct {
	gorm.Model
	CreditCards3 []CreditCard3 `gorm:"foreignKey:UserRefer"`
}

type CreditCard3 struct {
	gorm.Model
	Number    string
	UserRefer uint
}

// User9 GORM 通常使用拥有者的主键作为外键的值。 对于上面的例子，它是 User 的 ID 字段。
// 为 user 添加 credit card 时，GORM 会将 user 的 ID 字段保存到 credit card 的 UserID 字段。
// 同样的，您也可以使用标签 references 来更改它，例如：
type User9 struct {
	gorm.Model
	MemberNumber string
	CreditCards4 []CreditCard4 `gorm:"foreignKey:UserNumber;references:MemberNumber"`
}

type CreditCard4 struct {
	gorm.Model
	Number     string
	UserNumber string
}

// User10 多对多
// Many to Many 会在两个 model 中添加一张连接表。
// 例如，您的应用包含了 user 和 language，且一个 user 可以说多种 language，多个 user 也可以说一种 language。
// User 拥有并属于多种 language，`user_languages` 是连接表
type User10 struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

// User11 User 当使用 GORM 的 AutoMigrate 为 User 创建表时，GORM 会自动创建连接表
// 反向引用
// User 拥有并属于多种 language，`user_languages` 是连接表
type User11 struct {
	gorm.Model
	Languages1 []*Language1 `gorm:"many2many:user_languages;"`
}

type Language1 struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

// GetAllUsers 检索 User 列表并预加载 Language
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("Languages").Find(&users).Error
	return users, err
}

// GetAllLanguages 检索 Language 列表并预加载 User
func GetAllLanguages(db *gorm.DB) ([]Language, error) {
	var languages []Language
	err := db.Model(&Language{}).Preload("Users").Find(&languages).Error
	return languages, err
}

// User12 重写外键
// 对于 many2many 关系，连接表会同时拥有两个模型的外键
// 某些数据库只允许在唯一索引字段上创建外键，如果您在迁移时会创建外键，则需要指定 unique index 标签。
type User12 struct {
	gorm.Model
	Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;joinReferences:ProfileRefer"`
	Refer    uint      `gorm:"index:,unique"`
}

type Profile struct {
	gorm.Model
	Name      string
	UserRefer uint `gorm:"index:,unique"`
}

// User13 自引用 Many2Many
type User13 struct {
	gorm.Model
	Friends []*User `gorm:"many2many:user_friends"`
}

func main() {
	// --------------------------------------------------------------------------------  关联模式
	// 自动创建、更新
	// GORM在创建或更新记录时会自动地保存其关联和引用，主要使用upsert技术来更新现有关联的外键引用。
	// 在创建时自动保存关联
	// 当你创建一条新的记录时，GORM会自动保存它的关联数据。 这个过程包括向关联表插入数据以及维护外键引用。
	user := User{
		Name: "jinzhu",
		BillingAddress: Address{
			Address1: "Billing Address - Address 1",
		},
		ShippingAddress: Address{
			Address1: "Shipping Address - Address 1",
		},
		Emails: []Email{
			{
				Email: "jinzhu@example.com",
			},
			{
				Email: "jinzhu-2@example.com",
			},
		},
		Languages: []Language{
			{
				Name: "ZH",
			},
			{
				Name: "EN",
			},
		},
	}

	// 创建用户及其关联的地址、电子邮件和语言
	//db.Create(&usered)
	// BEGIN TRANSACTION;
	// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1"), ("Shipping Address - Address 1") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
	// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com"), (111, "jinzhu-2@example.com") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "languages" ("name") VALUES ('ZH'), ('EN') ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "user_languages" ("user_id","language_id") VALUES (111, 1), (111, 2) ON DUPLICATE KEY DO NOTHING;
	// COMMIT;
	//db.Save(&user)
	// 通过 FullSaveAssociations 来更新关联
	// 对于需要全面更新关联数据（不止外键）的情况，就应该使用 FullSaveAssociations 方法。
	// 更新用户并完全更新其所有关联
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	// SQL：完全更新地址、用户、电子邮件表，包括现有的关联记录
	// 不指定ID或者指定的ID不存在时会新增一条记录，否则是更新
	// user.BillingAddress = Address{Address1: "888"}
	user.BillingAddress = Address{Address1: "000", ID: user.BillingAddressID}
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

	// 跳过自动创建、更新
	// GORM 提供了在创建或更新操作过程中跳过自动保存关联的灵活性。 通过使用Select或者Omit方法可以允许您指定具体哪些字段在操作中被包含或者排除
	// 使用Select 来指定字段范围
	// Select方法可以让您模型中的哪些字段应该被保存 也就是说只有被选中的字段会被包含在SQL中
	// 当插入用户的时候仅包含“Name”字段
	db.Select("Name").Create(&user)

	// 使用Omit来排除字段或关联
	// 创建用户时跳过字段“BillingAddress”
	db.Omit("BillingAddress").Create(&user)
	// 创建用户时跳过全部关联关系
	db.Omit(clause.Associations).Create(&user)

	// 删除关联
	// GORM 能在删除主模型时使用Select方法来删除关联关系(一对一、一对多、多对多)。 在删除时维护好数据完整性并确保关联数据被妥当管理上，这项特性非常有用。
	// 可用Select来指定哪些关联应该随着主模型被删除
	//删除用户时，也删除用户的帐户
	db.Select("Account").Delete(&user)
	// 删除 user 时，也删除 user 的 Orders、CreditCards 关联记录
	db.Select("Orders", "CreditCards").Delete(&user)
	// 删除用户 时，也删除用户的所有一对一、一对多和多对多关联
	db.Select(clause.Associations).Delete(&user)
	// 删除多个用户时，也同步删除每个用户的账户信息。
	db.Select("Account").Delete(&users)

	// --关联模式
	// GORM 中的关联模式提供了多种辅助方法来处理模型之间的关系，为管理关联数据提供了高效的方式。
	// 要启动关联模式，您需要指定源模型和关系的字段名称。 源模型必须包含主键，并且关系的字段名称应与现有的关联字段相匹配。
	db.Model(&user).Association("Languages")
	// 检查error
	err := db.Model(&user).Association("Languages").Error
	fmt.Println(err)

	// --查询关联
	//检索关联记录，可以带有或不带有额外的条件。
	languages := ""
	err = db.Model(&user).Association("Languages").Find(&languages)
	if err != nil {
		return
	}
	// Find with conditions
	codes := []string{"zh-CN", "en-US", "ja-JP"}
	err = db.Model(&user).Where("code IN ?", codes).Association("Languages").Find(&languages)
	if err != nil {
		return
	}
	// 替换关联
	//用新的关联替换当前的关联。
	// Replace existing languages
	//db.Model(&user).Association("Languages").Replace([]Language{languageZH, languageEN})
	//db.Model(&user).Association("Languages").Replace(Language{Name: "DE"}, languageEN)
	// 删除关联
	//删除源模型和目标模型之间的关系，只删除引用。
	// Delete specific languages
	//db.Model(&user).Association("Languages").Delete([]Language{languageZH, languageEN})
	//db.Model(&user).Association("Languages").Delete(languageZH, languageEN)
	// 清空关联
	//删除源模型和关联之间的所有引用。
	// Clear all languages
	err = db.Model(&user).Association("Languages").Clear()
	if err != nil {
		return
	}
	// 关联计数
	//获取当前关联的数量，可以带有或不带有条件。
	// Count all languages
	db.Model(&user).Association("Languages").Count()
	// Count with conditions
	codes = []string{"zh-CN", "en-US", "ja-JP"}
	db.Model(&user).Where("code IN ?", codes).Association("Languages").Count()

	// User18 -----------------------------------------------------------------------------------------           预加载
	// Preload预加载
	// GORM允许使用 Preload通过多个SQL中来直接加载关系, 例如：
	type User18 struct {
		gorm.Model
		Username string
		Orders2  []Order2
	}
	// 查找 user 时预加载相关 Order
	db.Preload("Order2").Find(&users)
	// SELECT * FROM users;
	// SELECT * FROM orders WHERE user_id IN (1,2,3,4);
	db.Preload("Orders").Preload("Profile").Preload("Role").Find(&users)
	// SELECT * FROM users;
	// SELECT * FROM orders WHERE user_id IN (1,2,3,4); // has many
	// SELECT * FROM profiles WHERE user_id IN (1,2,3,4); // has one
	// SELECT * FROM roles WHERE id IN (4,5,6); // belongs to

	// -- Joins 预加载
	// Preload在单独的查询中加载关联数据，Join Preload将使用左连接加载关联数据。
	db.Joins("Company").Joins("Manager").Joins("Account").First(&user, 1)
	db.Joins("Company").Joins("Manager").Joins("Account").First(&user, "users.name = ?", "jinzhu")
	db.Joins("Company").Joins("Manager").Joins("Account").Find(&users, "users.id IN ?", []int{1, 2, 3, 4, 5})

	// 带条件的 Join
	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id` AND `Company`.`alive` = true;
	db.Joins("Company", db.Where(&Company{Alive: true})).Find(&users)
	// Join nested model
	// SELECT "users"."id","users"."created_at","users"."updated_at","users"."deleted_at","users"."name","users"."age","users"."birthday","users"."company_id","users"."manager_id","users"."active","Manager"."id" AS "Manager__id","Manager"."created_at" AS "Manager__created_at","Manager"."updated_at" AS "Manager__updated_at","Manager"."deleted_at" AS "Manager__deleted_at","Manager"."name" AS "Manager__name","Manager"."age" AS "Manager__age","Manager"."birthday" AS "Manager__birthday","Manager"."company_id" AS "Manager__company_id","Manager"."manager_id" AS "Manager__manager_id","Manager"."active" AS "Manager__active","Manager__Company"."id" AS "Manager__Company__id","Manager__Company"."name" AS "Manager__Company__name" FROM "users" LEFT JOIN "users" "Manager" ON "users"."manager_id" = "Manager"."id" AND "Manager"."deleted_at" IS NULL LEFT JOIN "companies" "Manager__Company" ON "Manager"."company_id" = "Manager__Company"."id" WHERE "users"."deleted_at" IS NULL
	db.Joins("Manager").Joins("Manager.Company").Find(&users)

	// User20 -- 预加载全部
	//关联可以与预加载类似，如创建/更新时的选择，您可以使用它来预加载所有关联
	db.Preload(clause.Associations).Find(&users)

	// -- 条件预加载
	// GORM允许与条件进行预加载关联，其工作原理类似于内联条件
	// Preload Orders with conditions
	db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
	// SELECT * FROM users;
	// SELECT * FROM orders WHERE user_id IN (1,2,3,4) AND state NOT IN ('canceled');
	db.Where("state = ?", "active").Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
	// SELECT * FROM users WHERE state = 'active';
	// SELECT * FROM orders WHERE user_id IN (1,2) AND state NOT IN ('canceled');

	// 自定义预加载 SQL
	//您可以通过传入func（db*gorm.db）*gorm来自定义预加载SQL
	// SELECT * FROM users;
	// SELECT * FROM orders WHERE user_id IN (1,2,3,4) order by orders.amount DESC;
	db.Preload("Orders", func(db *gorm.DB) *gorm.DB {
		return db.Order("orders.amount DESC")
	}).Find(&users)

	// 嵌套预加载
	// GORM支持嵌套预加载
	db.Preload("Orders.OrderItems.Product").Preload("CreditCard").Find(&users)

	// Customize Preload conditions for `Orders`
	// And GORM won't preload unmatched order's OrderItems then
	db.Preload("Orders", "state = ?", "paid").Preload("Orders.OrderItems").Find(&users)
}
