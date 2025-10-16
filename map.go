package main

import (
	"fmt"
	"sync"
	"time"
)

// map 集合是无序键值对集合。相比切片和数组，map 集合对索引的自定义程度更高，可以使用任意类型作为索引，也可以存储任意类型的数据。
// 但 map 集合中，存储的键值对的顺序是不确定的。当获取 map 集合中的值时，如果键不存在，则返回类型的零值。

func main12() {
	// -------------------------------------------------------------------------  声明map
	var m1 map[string]string
	fmt.Println("m1 length:", len(m1))

	m2 := make(map[string]string)
	fmt.Println("m2 length:", len(m2))
	fmt.Println("m2 =", m2)

	m3 := make(map[string]string, 10)
	fmt.Println("m3 length:", len(m3))
	fmt.Println("m3 =", m3)

	m4 := map[string]string{}
	fmt.Println("m4 length:", len(m4))
	fmt.Println("m4 =", m4)

	m5 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println("m5 length:", len(m5))
	fmt.Println("m5 =", m5)

	// --------------------------------------------------------------------------- 使用map 集合
	m := make(map[string]int, 10)
	m["1"] = int(1)
	m["2"] = int(2)
	m["3"] = int(3)
	m["4"] = int(4)
	m["5"] = int(5)
	m["6"] = int(6)

	// 获取元素
	value1 := m["1"]
	fmt.Println("m[\"1\"] =", value1)

	value1, exist := m["1"]
	fmt.Println("m[\"1\"] =", value1, ", exist =", exist)

	valueUnexist, exist := m["10"]
	fmt.Println("m[\"10\"] =", valueUnexist, ", exist =", exist)

	// 修改值
	fmt.Println("before modify, m[\"2\"] =", m["2"])
	m["2"] = 20
	fmt.Println("after modify, m[\"2\"] =", m["2"])

	// 获取map的长度
	fmt.Println("before add, len(m) =", len(m))
	m["10"] = 10
	fmt.Println("after add, len(m) =", len(m))

	// 遍历map集合main
	for key, value := range m {
		fmt.Println("iterate map, m[", key, "] =", value)
	}

	// 使用内置函数删除指定的key
	_, exist100 := m["10"]
	fmt.Println("before delete, exist 10: ", exist100)
	delete(m, "10")
	_, exist100 = m["10"]
	fmt.Println("after delete, exist 10: ", exist100)

	// 在遍历时，删除map中的key
	for key := range m {
		fmt.Println("iterate map, will delete key:", key)
		delete(m, key)
	}
	fmt.Println("m = ", m)

	// ------------------------------------------------------------------------------  map 传参
	// map 集合也是引用类型，和切片一样，将 map 集合作为参数传给函数或者赋值给另一个变量，它们都指向同一个底层数据结构，对 map 集合的修改，都会影响到原始实参。
	mm := make(map[string]int)
	mm["a"] = 1
	receiveMap(mm)
	fmt.Println("m =", mm)

	// ----------------------------------------------------------------------------   并发时使用 map 集合 (互斥锁方式)
	mc := make(map[string]int)
	var wg sync.WaitGroup // 定义一个等待组,  等待所有goroutine 执行完成
	var lock sync.Mutex   // 定义一个互斥锁, 在并发操作中保护共享资源,(映射mc) 防止多个 goroutine 同时修改导致数据竞争
	wg.Add(2)             // 向等待组添加2任务,后面启动2个go 等待这2个完成。

	go func() {
		for {
			lock.Lock()
			mc["a"]++
			lock.Unlock()
		}
	}()

	go func() {
		for {
			lock.Lock()
			mc["a"]++
			fmt.Println(mc["a"])
			lock.Unlock()
		}
	}()

	select { // select + time.After 实现超时控制
	case <-time.After(time.Second * 5): // 5 秒后向对应通道发一个值
		fmt.Println("timeout, stopping")
	}
	// sync.Map 适用于读多写少的场景，并且内存开销会比普通的 map 集合更大。
	// 所以碰到这种情况，更推荐使用普通的互斥锁来保证 map 集合的并发读写的线程安全性。

}

func receiveMap(param map[string]int) {
	fmt.Println("before modify, in receiveMap func: param[\"a\"] = ", param["a"])
	param["a"] = 2
	param["b"] = 3
}
