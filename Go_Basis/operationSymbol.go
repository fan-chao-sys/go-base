package Go_Basis

import "fmt"

//  运算操作符
// +  相加
// -  相减
// *  相乘
// /  相除
// %  求余
// ++ 相增
// -- 相减

func main8888() {
	// ---------------------------------------------------------------------------------------  算数运算符
	// 两个整数计算，它们计算之后的结果也是整数
	a, b := 1, 2
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	mod := a % b
	fmt.Println(sum, sub, mul, div, mod)

	// 自增自减只能单独声明存在,正确写法:  	   a++   a--
	// 错误的使用方式:  					   ++a   --a
	// 不可以自增时计算,也不能赋值，错误使用方式:  b := a++ + 1      c := a--

	// 当不同的数字类型混合计算时，必须先把它们转换成同一类型才可以计算：
	//a := 10 + 0.1
	//b := byte(1) + 1
	//fmt.Println(a, b)

	//sum := a + float64(b)
	//fmt.Println(sum)

	//sub := byte(a) - b
	//fmt.Println(sub)

	//mul := a * float64(b)
	//div := int(a) / b
	//fmt.Println(mul, div)

	// ---------------------------------------------------------------------------------------  关系运算符
	symbol2()

	// ---------------------------------------------------------------------------------------  逻辑运算符
	symbol3()

	// ---------------------------------------------------------------------------------------  位运算符
	symbol4()

}

// ---------------------------------------------------------------------------------------  赋值运算符
func main8765() {
	a, b := 1, 2
	var c int
	c = a + b
	fmt.Println("c = a + b, c =", c)

	plusAssignment(c, a)
	subAssignment(c, a)
	mulAssignment(c, a)
	divAssignment(c, a)
	modAssignment(c, a)
	leftMoveAssignment(c, a)
	rightMoveAssignment(c, a)
	andAssignment(c, a)
	orAssignment(c, a)
	norAssignment(c, a)
}

func plusAssignment(c, a int) {
	c += a // c = c + a
	fmt.Println("c += a, c =", c)
}

func subAssignment(c, a int) {
	c -= a // c = c - a
	fmt.Println("c -= a, c =", c)
}

func mulAssignment(c, a int) {
	c *= a // c = c * a
	fmt.Println("c *= a, c =", c)
}

func divAssignment(c, a int) {
	c /= a // c = c / a
	fmt.Println("c /= a, c =", c)
}

func modAssignment(c, a int) {
	c %= a // c = c % a
	fmt.Println("c %= a, c =", c)
}

// 左移运算符: 将左操作左移右操作数指定的位数,并将结果赋值给左操作数
func leftMoveAssignment(c, a int) {
	c <<= a // c = c << a
	fmt.Println("c <<= a, c =", c)
}

// 右移运算符: 将左操作右移右操作数指定的位数,并将结果赋值给左操作数
func rightMoveAssignment(c, a int) {
	c >>= a // c = c >> a
	fmt.Println("c >>= a, c =", c)
}

// 按位与赋值: 将左操作数与右操作数进行按位与运算，并将结果赋给左操作数
func andAssignment(c, a int) {
	c &= a // c = c & a
	fmt.Println("c &= a, c =", c)
}

// 按位异或赋值:将左操作数与右操作数进行按位异或运算，并将结果赋给左操作数
func orAssignment(c, a int) {
	c |= a // c = c | a
	fmt.Println("c |= a, c =", c)
}

// 按位或赋值:将左操作数与右操作数进行按位或运算，并将结果赋给左操作数
func norAssignment(c, a int) {
	c ^= a // c = c ^ a
	fmt.Println("c ^= a, c =", c)
}

// ---------------------------------------------------------------------------------------  其他运算符
func main66666() {
	a := 4
	var ptr *int
	fmt.Println(a)

	ptr = &a
	fmt.Printf("*ptr 为 %d\n", *ptr)
}

// ---------------------------------------------------------------------------------------  运算优先级
// 注：可以使用小括号，提高部分计算的优先级。也可以提高表达式的可读性。
// 优先级	运算符
//
//	4		* / % << >> &
//	3		+ - | ^
//	2		== != < <= > >=
//	1		&& ||
func main777() {
	var a int = 21
	var b int = 10
	var c int = 16
	var d int = 5
	var e int

	e = (a + b) * c / d // ( 31 * 16 ) / 5
	fmt.Printf("(a + b) * c / d 的值为 : %d\n", e)

	e = ((a + b) * c) / d // ( 31 * 16 ) / 5
	fmt.Printf("((a + b) * c) / d 的值为  : %d\n", e)

	e = (a + b) * (c / d) // 31 * (16/5)
	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n", e)

	// 21 + (160/5)
	e = a + (b*c)/d
	fmt.Printf("a + (b * c) / d 的值为  : %d\n", e)

	// 2 & 2 = 2; 2 * 3 = 6; 6 << 1 = 12; 3 + 4 = 7; 7 ^ 3 = 4;4 | 12 = 12
	f := 3 + 4 ^ 3 | 2&2*3<<1
	fmt.Println(f == 12)
}

func symbol2() {
	// 关系运算符结果只会是 bool 类型。
	a := 1
	b := 5
	fmt.Println(false)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
}

func symbol3() {
	a := true
	b := false

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!(a && b))
}

/*
*
位运算符
*/
func symbol4() {
	// & :  按位与   对两个操作数的每个位执行与操作，当且仅当两个位都为 1 时，结果位才为 1
	// | :  按位或   对两个操作数的每个位执行或操作，只要有一个位为 1，结果位就为 1
	// ^ :  按位异或 对两个操作数的每个位执行异或操作，当两个位不同时，结果位为 1
	// &^ : 按位清零 对于右操作数中为 1 的位，将左操作数中对应的位清零，其他位保持不变
	// << : 左移    将左操作数的所有位向左移动右操作数指定的位数，右侧空出的位用 0 填充
	// >> : 右移    将左操作数的所有位向右移动右操作数指定的位数，左侧空出的位用 0 填充（无符号数）或符号位填充（有符号数）

	// 注意：
	// 位运算符仅适用于整数类型（int、uint、int8 等），不能用于 float、string 等其他类型
	// 左移和右移操作中，右操作数必须是无符号整数

	fmt.Println(0 & 0)
	fmt.Println(0 | 0)
	fmt.Println(0 ^ 0)

	fmt.Println(0 & 1)
	fmt.Println(0 | 1)
	fmt.Println(0 ^ 1)

	fmt.Println(1 & 1)
	fmt.Println(1 | 1)
	fmt.Println(1 ^ 1)

	fmt.Println(1 & 0)
	fmt.Println(1 | 0)
	fmt.Println(1 ^ 0)
}
