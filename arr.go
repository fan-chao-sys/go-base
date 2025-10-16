package main

import "fmt"

// 数组
//	数据是具有相同类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的基础数据类型、自定义类型、指针以及其他数据结构。
//	数组元素通过索引（下标）来读取或修改，索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推，最后一个元素索引为数组的长度 - 1。

func main11() {

	// -------------------------------------------------------------------- 声明 创建 数组 初始化
	// 仅声明
	var a [5]int
	fmt.Println("a = ", a)

	var marr [2]map[string]string
	fmt.Println("marr = ", marr)
	// map的零值是nil，虽然打印出来是非空值，但真实的值是nil

	// 声明以及初始化
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	// 类型推导声明方式
	var c = [5]string{"c1", "c2", "c3", "c4", "c5"}
	fmt.Println("c = ", c)
	d := [3]int{3, 2, 1}
	fmt.Println("d = ", d)

	// 使用 ... 代替数组长度
	autoLen := [...]string{"auto1", "auto2", "auto3"}
	fmt.Println("autoLen = ", autoLen)

	// 声明时初始化指定下标的元素值
	positionInit := [5]string{1: "position1", 3: "position3"}
	fmt.Println("positionInit = ", positionInit)

	// ------------------------------------------------------------------- 访问数组
	ac := [5]int{5, 4, 3, 2, 1}

	// 方式1，使用下标读取数据
	element := ac[2]
	fmt.Println("element = ", element)

	// 方式2，使用range遍历
	for i, v := range ac {
		fmt.Println("index = ", i, "value = ", v)
	}
	for i := range ac {
		fmt.Println("only index, index = ", i)
	}

	// 读取数组长度
	fmt.Println("len(a) = ", len(ac))
	// 使用下标，for循环遍历数组
	for i := 0; i < len(ac); i++ {
		fmt.Println("use len(), index = ", i, "value = ", a[i])
	}

	// ------------------------------------------------------------------  多维数组
	// 声明方式，不限制多维数组的嵌套层数：
	// 二维数组
	ar := [3][2]int{
		{0, 1},
		{2, 3},
		{4, 5},
	}
	fmt.Println("a = ", ar)

	// 三维数组
	br := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}
	fmt.Println("b = ", br)

	// 也可以省略各个位置的初始化,在后续代码中赋值
	cr := [3][3][3]int{}
	cr[2][2][1] = 5
	cr[1][2][1] = 4
	fmt.Println("c = ", cr)

	// 访问多维数组与访问普通数组的方式一致：
	// 三维数组
	arr := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}

	layer1 := arr[0]
	layer2 := arr[0][1]
	element = arr[0][1][1]
	fmt.Println(layer1)
	fmt.Println(layer2)
	fmt.Println(element)

	// 多维数组遍历时，需要使用嵌套for循环遍历
	for i, v := range arr {
		fmt.Println("index = ", i, "value = ", v)
		for j, inner := range v {
			fmt.Println("inner, index = ", j, "value = ", inner)
		}
	}

	// --------------------------------------------------------------------------------  数组作参数
	// 数组部分特性类似基础数据类型，当数组作为参数传递时，在函数中并不能改变外部实参的值。
	// 如果想要修改外部实参的值，需要把数组的指针作为参数传递给函数。
	var carr = [5]*Customs{
		{6},
		{7},
		{8},
		{9},
		{10},
	}
	ao := [5]int{5, 4, 3, 2, 1}
	fmt.Println("before all, a = ", ao)
	printFuncParamPointer(carr)
	receiveArray(ao)
	fmt.Println("after receiveArray, a = ", ao)

	receiveArrayPointer(&ao)
	fmt.Println("after receiveArrayPointer, a = ", ao)
}

type Customs struct {
	i int
}

func receiveArray(param [5]int) {
	fmt.Println("in receiveArray func, before modify, param = ", param)
	param[1] = -5
	fmt.Println("in receiveArray func, after modify, param = ", param)
}

func receiveArrayPointer(param *[5]int) {
	fmt.Println("in receiveArrayPointer func, before modify, param = ", param)
	param[1] = -5
	fmt.Println("in receiveArrayPointer func, after modify, param = ", param)
}

func printFuncParamPointer(param [5]*Customs) {
	for i := range param {
		fmt.Printf("in printFuncParamPointer func, param[%d] = %p, value = %v \n", i, &param[i], *param[i])
	}

}
