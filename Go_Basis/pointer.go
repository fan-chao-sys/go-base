package Go_Basis

import (
	"fmt"
	"unsafe"
)

// 声明一个指针类型变量需使用星号 * 标识
// var <name> *<type>
var p1 *int
var p2 *string

func main4() {

	// 初始化指针必须通过另外一个变量， 如果没有赋值
	// p = &<var name>
	// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针,因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
	w := 1
	s := "Hello"
	p1 = &w
	p2 = &s
	p3 := &p2

	// 也可用一个结构体实例或者变量直接声明并且赋值给一个指针
	// p := &<struct type>{}

	// 还可以获取指针的指针：
	var p **string

	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(*p1)
	fmt.Println(*p2)
	fmt.Println(**p3)

	// 通过指针修改原始变量的值
	var p1 *int
	i := 1
	p1 = &i
	fmt.Println(*p1 == i)
	// 指针的引用内存地址 可以赋值 字面量
	*p1 = 2
	fmt.Println(i)

	// 修改指针指向的值
	abc := 2                 // 局部变量,整数值类型2 abc
	var pa *int              // 指针变量 pa
	fmt.Println("abc", &abc) // 打印abc内存地址
	pa = &abc                // abc内存地址赋值给 指针 pa
	fmt.Println(pa, &abc)    // 打印 pa指针, abc内存地址

	var pp **int // 定义指针的指针变量 pp
	pp = &pa     // 赋值 pa变量的第一次指针内存地址给 pp
	fmt.Println(pp, pa)
	**pp = 3                 // pp的指针的指针二次回溯地址的值，赋值为3，等于 abc改为3
	fmt.Println(pp, *pp, pa) // 打印，pp二次指针地址，pp一次指针地址，pa指针地址
	fmt.Println(**pp, *pa)   // 打印 pp二次指针最终引用地址值， pa的引用地址值
	fmt.Println(abc, &abc)   // 打印 abc的值,abc内存地址

	// 总结: 普通变量数据存在于 内存地址 及 内存数值(真实数值)上 / 指针变量 的内存地址和内存数值都是引用其他变量的内存地址。

	// ------------------------------------------------------------ 指针偏移访问 unsafe.Pointer 和 uintptr
	// 解释：指针类型不能直接进行计算，需要通过 unsafe.Pointer转换后,在转 uintptr整数类型，才能支持指针类型计算
	// *T <---> unsafe.Pointer <---> uintptr

	// 1.把指针转换成 unsafe.Pointer:
	var ps *int
	var as int = 1
	ps = &as
	up1 := unsafe.Pointer(ps)
	up2 := unsafe.Pointer(&as)
	fmt.Println("u1,u2", up1, up2)

	// 2.把unsafe.Pointer 转成 uintptr
	fmt.Println("uintptr1", uintptr(up1))

	// 注意，这个操作非常危险，并且结果不可控，在一般情况下是不需要进行这种操作。
	ac := "Hello, world!"
	upA := uintptr(unsafe.Pointer(&ac))
	upA += 1

	cte := (*uint8)(unsafe.Pointer(upA))
	fmt.Println("*c", *cte)

}
