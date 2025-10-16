package main

// 变量的作用域
import (
	"fmt"
	"time"
)

// -------------------------------------------------------------------------------- 局部变量
// 在函数内声明的变量，作用域范围只在函数体内。
// 同理，函数的参数和返回值也是局部变量。 parameter、res、decVar 都是局部变量，仅在当前函数中有效。
func localVariable(parameter int) (res int) {
	var _ int
	return
}

// 比较特殊的是 if 语句、for 语句、switch 语句、select 语句、匿名代码块中声明的变量，仅在小代码块内有效。
func main10() {
	var a int
	if b := 1; b == 0 {
		fmt.Println("b == 0")
	} else {
		c := 2
		fmt.Println("declare c = ", c)
		fmt.Println("b == 1")
	}

	switch d := 3; d {
	case 1:
		e := 4
		fmt.Println("declare e = ", e)
		fmt.Println("d == 1")
	case 3:
		f := 4
		fmt.Println("declare f = ", f)
		fmt.Println("d == 3")
	}

	for i := 0; i < 1; i++ {
		forA := 1
		fmt.Println("forA = ", forA)
	}

	select {
	case <-time.After(time.Second):
		selectA := 1
		fmt.Println("selectA = ", selectA)
	}

	// 匿名代码块
	{
		blockA := 1
		fmt.Println("blockA = ", blockA)
	}
	fmt.Println("a = ", a)
}

// ----------------------------------------------------------------------- 全局变量
// 全局变量函数外声明，全局变量作用域可以是当前整个包甚至外部包（公开的全局变量）。
// 当全局变量 和 局部变量 重名时，函数内会使用局部变量，超出局部变量作用域之后，才会重新使用全局变量。

var av int

func main100() {
	{
		fmt.Println("global variable, a = ", av)
		av = 3
		fmt.Println("global variable, a = ", av)

		av := 10
		fmt.Println("local variable, a = ", av)
		av--
		fmt.Println("local variable, a = ", av)
	}
	fmt.Println("global variable, a = ", av)
}

// 这种优先使用作用域更小的变量的规则，同样适用于局部变量：
func main1000() {
	var b int = 4
	fmt.Println("local variable, b = ", b)
	if b := 3; b == 3 {
		fmt.Println("if statement, b = ", b)
		b--
		fmt.Println("if statement, b = ", b)
	}
	fmt.Println("local variable, b = ", b)
}

// 实际代码使用中，经常会有各种方法返回 error，error 会赋值给 err 变量。
