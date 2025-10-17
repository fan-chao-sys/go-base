package Go语言基础

import (
	"fmt"
	"reflect"
	"time"
)

// range 迭代
// range 关键字用于 for 循环迭代字符串(string)、数组(array)、切片(slice)、通道(channel)、映射集合(map)中元素

// ----------------------------------------------------------------------------------------------------- 字符串迭代
// string 类型是一个比较特殊的类型，可以与 rune 切片类型、byte 切片类型相互转换，同时还可使用 range 关键字来遍历一个字符串。

func main121() {
	// 字符串都按 Unicode 编码
	str1 := "abc123"
	// 仅使用 range 获取下标索引
	// 第一 for 循环中，遍历了变量 str1，它有六个字符，而且这些字符都可以使用一个 byte 表示，所以循环了六次才退出循环。
	for index := range str1 {
		fmt.Printf("str1 -- index:%d, value:%d\n", index, str1[index])
	}

	str2 := "测试中文"
	// 第二 for 循环，遍历str2，四个中文字符，按Unicode 编码标准，很显然不能仅仅被四个 byte 表示，但它还是只循环了 4 次，恰好是中文字符的长度。
	// 那么就是遍历字符串时，实际是遍历从字符串转换来的 rune 切片，恰好字符串转换成 byte 切片 和 字符串转换成 rune 切片之后的长度相同，
	for index := range str2 {
		fmt.Printf("str2 -- index:%d, value:%d\n", index, str2[index])
	}
	fmt.Printf("len(str2) = %d\n", len(str2))
	runesFromStr2 := []rune(str2)
	bytesFromStr2 := []byte(str2)
	fmt.Printf("len(runesFromStr2) = %d\n", len(runesFromStr2))
	fmt.Printf("len(bytesFromStr2) = %d\n", len(bytesFromStr2))

	// 使用 range 获取下标和下标位置的字符
	// 当直接使用下标取字符串某个下标位置上的值时，取出来的是 byte 值。
	// 但当使用 range 关键字直接获取某个下标位置值时，取出的是一个完整 rune 类型值。
	str3 := "a1中文"
	for index, value := range str3 {
		fmt.Printf("str1 -- index:%d, index value:%d\n", index, str3[index])
		fmt.Printf("str1 -- index:%d, range value:%d\n", index, value)
	}

	// --------------------------------------------------------------------------------------------- 数组与切片迭代
	// 遍历一维数组与切片
	array := [...]int{1, 2, 3}
	slice := []int{4, 5, 6}
	// 方法1：只拿到数组的下标索引
	for index := range array {
		fmt.Printf("array -- index=%d value=%d \n", index, array[index])
	}
	for index := range slice {
		fmt.Printf("slice -- index=%d value=%d \n", index, slice[index])
	}
	fmt.Println()

	// 方法2：同时拿到数组的下标索引和对应的值
	for index, value := range array {
		fmt.Printf("array -- index=%d index value=%d \n", index, array[index])
		fmt.Printf("array -- index=%d range value=%d \n", index, value)
	}
	for index, value := range slice {
		fmt.Printf("slice -- index=%d index value=%d \n", index, slice[index])
		fmt.Printf("slice -- index=%d range value=%d \n", index, value)
	}
	fmt.Println()

	// 遍历二维数组与切片
	arrays := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	slices := [][]int{{1, 2}, {3}}
	// 只拿到行的索引
	for index := range arrays {
		// array[index]类型是一维数组
		fmt.Println(reflect.TypeOf(arrays[index]))
		fmt.Printf("array -- index=%d, value=%v\n", index, arrays[index])
	}

	for index := range slices {
		// slice[index]类型是一维数组
		fmt.Println(reflect.TypeOf(slices[index]))
		fmt.Printf("slice -- index=%d, value=%v\n", index, slices[index])
	}

	// 拿行索引和该行的数据
	// 切片与数组相比，特殊的地方就在于其长度可变，所以构成二维时，切片中元素的数量可以随意设置，而数组是定长的。
	// 使用 range 迭代时，两者体验完全一致。
	fmt.Println("print array element")
	for rowIndex, rowValue := range array {
		fmt.Println(rowIndex, reflect.TypeOf(rowValue), rowValue)
	}

	fmt.Println("print array slice")
	for rowIndex, rowValue := range slice {
		fmt.Println(rowIndex, reflect.TypeOf(rowValue), rowValue)
	}

	// 双重遍历，拿到每个元素的值
	for rowIndex, rowValue := range arrays {
		for colIndex, colValue := range rowValue {
			fmt.Printf("array[%d][%d]=%d ", rowIndex, colIndex, colValue)
		}
		fmt.Println()
	}
	for rowIndex, rowValue := range slices {
		for colIndex, colValue := range rowValue {
			fmt.Printf("slice[%d][%d]=%d ", rowIndex, colIndex, colValue)
		}
		fmt.Println()
	}

	// -------------------------------------------------------------------------------------------------- 通道迭代
	// 通道除了使用 for 循环配合 select 关键字获取数据以外，也可使用 for 循环配合 range 关键字获取数据。
	// 因为通道结构的特殊性，当使用 range 遍历通道时，只给一个迭代变量赋值，而不像数组或字符串一样能够使用 index 索引。
	// 当通道被关闭时，在 range 关键字迭代完通道中所有值后，循环就会自动退出。
	ch := make(chan int, 10)
	go addData(ch)
	for i := range ch {
		fmt.Println(i)
	}

	// ----------------------------------------------------------------------------------------------- 映射集合map迭代
	// 使用 range 关键字迭代映射集合时，一种是拿到 key，一种是拿到 key 和 value，并且 range 关键字在迭代映射集合时，其中的 key 是乱序的。
	hash6 := map[string]int{
		"a": 1,
		"f": 2,
		"z": 3,
		"c": 4,
	}

	for key := range hash6 {
		fmt.Printf("key=%s, value=%d\n", key, hash6[key])
	}

	for key, value := range hash6 {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
}

func addData(ch chan int) {
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
