package main

// ------------------------------------------------------------ const 常量
//	常量值在编译期确定，所以常量定义必须赋值，并不能方法返回值为常量赋值。
//	常量被定义后，其值不能再被修改。
//	常量（包括全局常量和局部常量）被定义后可以不使用。
//	常量定义方式与变量定义的方式基本相同，只是 var 关键字被更换成了 const：
// 注：Go 中，常量只能使用基本数据类型，即数字、字符串和布尔类型。不能使用复杂的数据结构，比如切片、数组、map、指针和结构体等。如果使用了非基本数据类型，会在编译期报错。

// const <name> <type> = <value>
const a int = 1

// 直接推导出来，不需要声明
const b = "test"

// const <name3>, <name4>, ... = <value3>, <value4>, ...
const c, d = 2, "hello"

// const <name5>, <name6>, ... <type> = <value5>, <value6>, ...
const e, f bool = true, false

// 声明多个时，可用括号包裹，此模式不限制声明次数
const (
	h    byte = 3
	i         = "value"
	j, k      = "v", 4
	l, m      = 5, false
	n         = 6
)

// ------------------------------------------------------------- const 枚举
// 枚举的本质就是一系列的常量(大写)
const (
	Male   = "Male"
	Female = "Female"
)

// Gender 除了直接定义值以外，还会使用类型别名，让常量定义的枚举类型的作用显得更直观，比如
type Gender string

const (
	Male1   Gender = "Male"
	Female1 Gender = "Female"
)

// 当此枚举作为参数传递时，会使用 Gender 作为参数类型，而不是基础类型 string，比如
func methodConstant(gender Gender) {}

// 并且使用了类型别名后，还可以为这个别名类型声明自定义方法：
func (g *Gender) String() string {
	switch *g {
	case Male1:
		return "Male"
	case Female1:
		return "Female"
	default:
		return "Unknown"
	}
}

func (g *Gender) IsMale() bool {
	return *g == Male
}

// ConnState ---------------------------------------------------------------  iota 关键字
// 除了上面的别名类型来声明枚举类型以外，还可以使用 iota 关键字，来自动为常量赋值。
// iota 辅助声明枚举之前，先了解一下 iota 关键字的使用。
// 标准库中的使用:
type ConnState int

const (
	StateNew ConnState = iota
	StateActive
	StateIdle
	StateHijacked
	StateClosed
)

type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// iota 方便定义常量的关键字。
//
//	iota 独立作用于每个 const 定义组，就是上面看到的 const ( ``// ``code... ) 结构。
//	并且每个 const 语句算作是一个 const 定义组。
//	如果 iota 定义在 const 定义组中的第 n 行，那 iota 值为 n-1.所以一定要注意 iota 出现在定义组中的第几行，而不是当前代码中它第几次出现。
const pre int = 1
const aa int = iota
const (
	bb int = iota
	cc
	dd
	ee
	ff = 2
	gg = iota
	hh
	ii
)

// Genders 使用 iota 关键字就是为了方便我们定义常量的值。
//
//	并且当这些枚举值仅作为判断条件使用时，修改非常方便，只需要的其分组增删即可。
////	注：iota 仅能与 const 关键字配合使用。
//type Genders byte
//
//const (
//	Maleed Genders = iota
//	Femaleed
//)
