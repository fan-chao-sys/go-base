package main

import "fmt"

func main2() {
	// ----------------------------------------------------- 进制
	// 二进制
	var f uint8 = 0b1111
	var g uint8 = 0b1111
	// 八进制
	var c uint8 = 017
	var d uint8 = 0o17
	var e uint8 = 0o17
	// 十进制
	var h uint8 = 15
	// 十六进制
	var a uint8 = 0xF
	var b uint8 = 0xf
	fmt.Println(a == b)
	fmt.Println(b == c)
	fmt.Println(c == d)
	fmt.Println(d == e)
	fmt.Println(e == f)
	fmt.Println(f == g)
	fmt.Println(g == h)

	// ---------------------------------------------------- 浮点数
	var float1 float32 = 10
	float2 := 10.0
	fmt.Println(float64(float1) == float2)

	// ---------------------------------------------------- complex复数
	num1 := 3 + 4i
	num2 := 1 + 2i
	sum := num1 + num2
	sub := num1 - num2
	mul := num1 * num2
	div := num1 / num2
	fmt.Printf("加法结果: %v\n", sum)
	fmt.Printf("减法结果: %v\n", sub)
	fmt.Printf("乘法结果: %v\n", mul)
	fmt.Printf("除法结果: %v\n", div)

	var c1 complex64
	c1 = 1.10 + 0.1i
	c2 := 1.10 + 0.1i
	c3 := complex(1.10, 0.1) // c2与c3是等价的
	x := real(c2)
	y := imag(c2)
	fmt.Printf("c1: %v\n", c1)
	fmt.Printf("c2: %v\n", c2)
	fmt.Printf("c3: %v\n", c3)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)

	// ------------------------------------------------- byte/ uint8
	// 字符串可以直接被转换成 []byte（byte 切片）。
	var s123 string = "Hello, world!"
	var bytes []byte = []byte(s123)
	fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
	var ss string = string(bytes)
	fmt.Println(ss)

	// []byte 也可以直接转换成 string
	var bytes1 []byte = []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
	fmt.Println(string(bytes1))

	// ------------------------------------------------- rune/ int32
	// Unicode 码点
	var r1 rune = 'a'
	var r2 rune = '世'
	fmt.Println(r1, r2)
	// 字符串可以直接转换成 []rune（rune 切片）
	var s string = "abc，你好，世界！"
	var runes []rune = []rune(s)
	fmt.Println(runes)
	fmt.Println(string(runes))
	fmt.Println(len(runes))

	// ------------------------------------------------- string
	// 双引号定义风格
	var s1 string = "Hello\nworld!\n"
	// 反引号定义风格（包括换行等都是整体）
	var s2 string = `Hello
world!
`
	fmt.Println(s1 == s2)

	// ------------------------------------------------- byte、rune 与 string 之间的联系
	var sssss string = "Go语言"
	var bytessss []byte = []byte(sssss)
	var runessss []rune = []rune(sssss)
	fmt.Println("string length: ", len(sssss))
	fmt.Println("bytes length: ", len(bytessss))
	fmt.Println("runes length: ", len(runessss))

	var sw string = "Go语言"
	var bytesw []byte = []byte(sw)
	var runesw []rune = []rune(sw)

	fmt.Println("string sub: ", sw[0:7])
	fmt.Println("bytes sub: ", string(bytesw[0:7]))
	fmt.Println("runes sub: ", string(runesw[0:3]))

	// -------------------------------------------------- bool 布尔类型(默认 false)
	var flag bool = true
	fmt.Println(flag)

	// -------------------------------------------------- 零值
	// 	每种类型都有零值，一个类型零值  =  默认值，当一个类型的变量没有被初始化时，其值就是默认值。
	//	数字类型的零值都是 0。
	//	字符串类型的零值是空字符串。
	//	布尔类型的零值是 false。
}
