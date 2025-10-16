package main

import (
	"fmt"
	"unsafe"
)

// 切片
// 切片(Slice)并不是数组或者数组指针，而是数组的一个引用，
// 切片本身是一个标准库中实现的一个特殊的结构体，这个结构体中有三个属性，分别代表数组指针、长度、容量。
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func slice1() {
	// ------------------------------------------------------------------------------- 声明与初始化切片
	// 切片的声明方式与声明数组的方式非常相似，与数组相比，切片不用声明长度:
	// 方式1，声明并初始化一个空的切片
	var s1s []int = []int{}
	// 方式2，类型推导，并初始化一个空的切片
	var s2s = []int{}
	// 方式3，与方式2等价
	s3s := []int{}
	// 方式4，与方式1、2、3 等价，可以在大括号中定义切片初始元素
	s4 := []int{1, 2, 3, 4}
	// 方式5，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为0
	s5 := make([]int, 0)
	// 方式6，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为2，指定容量参数4
	s6 := make([]int, 2, 4)
	// 方式7，引用一个数组，初始化切片
	arr := [5]int{6, 5, 4, 3, 2}
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := arr[2:]
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s8 := arr[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s9 := arr[:2]
	fmt.Print(s1s)
	fmt.Print(s2s)
	fmt.Print(s3s)
	fmt.Print(s4)
	fmt.Print(s5)
	fmt.Print(s6)
	fmt.Print(s7)
	fmt.Print(s8)
	fmt.Print(s9)

	// 当切片是基于同一个数组指针创建出来时，修改数组中的值时，同样会影响到这些切片。
	aq := [5]int{6, 5, 4, 3, 2}
	// 从数组下标2开始，直到数组的最后一个元素
	s71 := aq[2:]
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s81 := aq[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s91 := aq[:2]
	fmt.Println(s71)
	fmt.Println(s81)
	fmt.Println(s91)

	// ------------------------------------------------------------------ 访问切片
	s1 := []int{5, 4, 3, 2, 1}
	// 下标访问切片
	e1 := s1[0]
	e2 := s1[1]
	e3 := s1[2]
	fmt.Println(s1)
	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)

	// 向指定位置赋值
	s1[0] = 10
	s1[1] = 9
	s1[2] = 8
	fmt.Println(s1)

	// range迭代访问切片
	for i, v := range s1 {
		fmt.Printf("before modify, s1[%d] = %d\n", i, v)
	}

	// 切片还可以使用 len() 和 cap() 函数访问切片的长度和容量。
	// 长度表示切片可以访问到底层数组的数据范围。
	// 容量表示切片引用的底层数组的长度。
	// 当切片是 nil 时，len() 和 cap() 函数获取到值都是 0。
	// 切片的长度小于等于切片的容量。
	var nilSlice []int
	fmt.Println("nilSlice length:", len(nilSlice))
	fmt.Println("nilSlice capacity:", len(nilSlice))

	s2 := []int{9, 8, 7, 6, 5}
	fmt.Println("s2 length: ", len(s2))
	fmt.Println("s2 capacity: ", cap(s2))

	//  切片添加元素
	// 切片是变长的，可以向切片追加新的元素，可以使用内置的 append() 向切片追加元素。
	// 内置函数 append() 只有切片类型可以使用，第一个参数必须是切片类型，后面追加的元素参数是变长类型，一次可以追加多个元素到切片。并且每次 append() 都会返回一个新的切片引用。
	s3 := []int{}
	fmt.Println("s3 = ", s3)

	// append函数追加元素
	s3 = append(s3)
	s3 = append(s3, 1)
	s3 = append(s3, 2, 3)
	fmt.Println("s3 = ", s3)

	// 除了使用 append() 函数向切片追加元素以外，还可以使用 append() 向指定位置添加元素，以及移除指定位置的元素。
	// 向指定位置添加元素的代码
	s4 = []int{1, 2, 4, 5}
	s4 = append(s4[:2], append([]int{3}, s4[2:]...)...)
	fmt.Println("s4 = ", s4)

	// 移除指定位置元素代码
	s5 = []int{1, 2, 3, 5, 4}
	s5 = append(s5[:3], s5[4:]...)
	fmt.Println("s5 = ", s5)

	// 复制切片
	// 可以使用内置函数 copy() 把某个切片中的所有元素复制到另一个切片，复制的长度是它们中最短的切片长度。
	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 3)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

	copy(dst1, src1)
	copy(dst2, src2)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

	// 切片底层原理
	// 切片类型实际上是比较特殊的指针类型，当声明一个切片类型时，就是声明了一个指针。
	// 这个指针指向的切片结构体，切片结构体中记录的三个属性：数组指针、长度、容量。这几个属性在创建一个切片时就定义好，并且在之后都不能再被修改。
	s := make([]int, 3, 6)
	fmt.Println("s length:", len(s))
	fmt.Println("s capacity:", cap(s))
	fmt.Println("initial, s = ", s)
	s[1] = 2
	fmt.Println("set position 1, s = ", s)
	modifySlice(s)
	fmt.Println("after modifySlice, s = ", s)

	// 在不使用 append() 函数的情况下，在函数内部对切片的修改，都会影响到原始实例。
	// 使用 append()函数时，需要分两种情况：
	// 当没有触发切片扩容时：
	s = make([]int, 3, 6)
	fmt.Println("initial, s =", s)
	s[1] = 2
	fmt.Println("after set position 1, s =", s)
	s2 = append(s, 4)
	fmt.Println("after append, s2 length:", len(s2))
	fmt.Println("after append, s2 capacity:", cap(s2))
	fmt.Println("after append, s =", s)
	fmt.Println("after append, s2 =", s2)
	s[0] = 1024
	fmt.Println("after set position 0, s =", s)
	fmt.Println("after set position 0, s2 =", s2)
	appendInFunc(s)
	fmt.Println("after append in func, s =", s)
	fmt.Println("after append in func, s2 =", s2)

	// 当使用 append() 函数之后。
	// 原来的切片引用，长度和容量不变，新追加的值超过切片可访问范围，访问不到新追加的值。
	// 新的切片引用，与原始切片引用相比，长度加一，容量不变，可以访问到新追加的值。
	// 在方法内，使用原始切片作为参数，使用 append() 函数追加元素后，同样会创建一个新的切片引用，新追加的值会覆盖之前的值。
	// 并且修改这个切片，其修改同样会反应到原始切片以及新的切片引用上。
	// 当 append() 函数触发扩容时：
	s = make([]int, 2, 2)
	fmt.Println("initial, s =", s)
	s2 = append(s, 4)
	fmt.Println("after append, s length:", len(s))
	fmt.Println("after append, s capacity:", cap(s))
	fmt.Println("after append, s2 length:", len(s2))
	fmt.Println("after append, s2 capacity:", cap(s2))
	fmt.Println("after append, s =", s)
	fmt.Println("after append, s2 =", s2)
	s[0] = 1024
	fmt.Println("after set position 0, s =", s)
	fmt.Println("after set position 0, s2 =", s2)
	appendInFunc(s2)
	fmt.Println("after append in func, s2 =", s2)

	// 当 append() 函数触发扩容后，实际上是新创建了一个数组实例，把原来的数组中的数据复制到了新数组中，然后创建一个新的切片实例并返回。
	// 这时原始切片中持有的数组指针指向的数组与新切片引用中的数组指针指向的数组是两个不同的数组，修改并不会相互影响。
	// 切片触发扩容前，切片一直共用相同的数组；
	// 切片触发扩容后，会创建新的数组，并复制这些数据；
	// 切片本身是一个特殊的指针，go 针对切片类型添加了一些语法糖，方便使用
}

func modifySlice(param []int) {
	param[0] = 1024
}

func appendInFunc(param []int) {
	param = append(param, 1022)
	fmt.Println("in func, param =", param)
	param[2] = 512
	fmt.Println("set position 2 in func, param =", param)
}

func appendInFunc2(param []int) {
	param1 := append(param, 511)
	param2 := append(param1, 512)
	fmt.Println("in func, param1 =", param1)
	param2[2] = 500
	fmt.Println("set position 2 in func, param2 =", param2)
}
