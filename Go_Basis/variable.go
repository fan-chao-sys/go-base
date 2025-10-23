package Go_Basis

import "fmt"

// ---------------------------------- 全局变量（允许声明后不用）
// var <name> <type> = <value> 完整声明形式
var a100 int = 100

// 仅声明未赋值
var b100 string

// 不声明类型，自动推导变量类型
var flag = true

// 对全局变量分组声明,声明多个，用括号包裹，此方式不限制声明次数
var s1 string = "Hello"
var zero int
var b1 = true

var (
	i100 int = 123
	b2   bool
	s2   = "test"
)

func main3() {

	// --------------------------------------- 局部变量声明
	// var <name> <type> = <value> 完整声明形式
	var name string = "熊"
	fmt.Println(name)

	// 仅声明，为类型默认零值：
	var a int
	fmt.Println(a)

	// 无需关键字 var，也无需声明类型，Go 通过字面量或表达式推导此变量类型：
	age := 23
	fmt.Println(age)

}

// --------------------------------------------- 局部声明变量
// 方式4，直接返回值中声明，方法一开始声明的变量
func method1() {
	// 类型推导，用得最多
	a := 1
	// 完整变量声明写法
	var b int = 2
	// 仅声明变量，但是不赋值，
	var c int
	fmt.Println(a, b, c)
}

func method2() (a int, b string) {
	// 这种方式必须声明return关键字，并同样不需使用，并也不用变量赋值
	return 1, "test"
}

func method3() (a int, b string) {
	a = 1
	b = "test"
	return
}

func method4() (a int, b string) {
	return
}

// ---------------------------------------------------  多变量定义
// 全局变量和局部变量都支持一次声明和定义多个变量。
// 全局变量声明方式：
// var <name1>, <name2>, ... <type> = <value1>, <value2>, ...
var address, phone string = "北京", "18201000891"

// var <name1>, <name2>, ... <type>
var sign, flags bool

// var <name1>, <name2>, ... = <value1>, <value2>, ...
var foot, ages = true, 23

// 局部变量声明方式：
// <name1>, <name2>, ... := <value1>, <value2>, ...
var a1, b12, c1 int = 1, 2, 3
var e1, f1, g1 int
var h1, i1, j1 = 1, 2, "test"

func method() {
	var k, l, m int = 1, 2, 3
	var n, o, p int
	q, r, s := 1, 2, "test"
	fmt.Println(k, l, m, n, o, p, q, r, s)
}
