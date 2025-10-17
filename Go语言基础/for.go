package Go语言基础

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func loopFor() {
	// --------------------------------------------------------------------------- for 循环

	// 方式1
	for i := 0; i < 10; i++ {
		fmt.Println("方式1，第", i+1, "次循环")
	}

	// 方式2   仅声明条件判断语句
	b := 1
	for b < 10 {
		fmt.Println("方式2，第", b, "次循环")
	}

	// 方式3，无限循环
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2)) // 创建带截止时间的上下文
	// context.Background() 是一个默认的、非 nil 的空上下文，通常作为其他上下文的父上下文。
	// context.WithDeadline 会创建一个新的上下文，该上下文会在指定的截止时间（这里是当前时间加上 2 秒）到期。当截止时间到期，或者父上下文被取消时，这个新上下文就会被取消。
	// 第二个返回值是一个 CancelFunc（用于主动取消上下文），这里用 _ 忽略了，因为我们主要依赖截止时间来触发上下文的取消。
	var started bool
	var stopped atomic.Bool // atomic.Bool类型变量，atomic 包提供对基本数据类型原子操作，确保多线程下对该变量操作线程安全，标记是否应该停止主循环。
	for {
		if !started {
			started = true
			go func() { // go 启动 ** goroutine（轻量级线程）** 的语法。go 关键字会创建一个新的 goroutine 并立即执行后面的函数。
				// 为在后台异步执行一段逻辑，不会阻塞当前（主）goroutine 的执行，让程序可以并发地做多个事情。
				for {
					select { // 用于监听多个通道（channel）操作的特殊语法，它会 “等待” 多个 case 中任意一个可以执行的情况，然后执行对应的逻辑。
					// 在这个代码片段里，select 只监听了一个 case：case <-ctx.Done():。
					// ctx.Done() 是一个通道，当上下文（Context）被取消（比如超时、手动取消等）时，这个通道会被关闭，此时 case <-ctx.Done(): 就会被触发。
					case <-ctx.Done():
						fmt.Println("ctx done")
						stopped.Store(true)
						return
					}
				}
			}()
		}
		fmt.Println("main")
		if stopped.Load() {
			break
		}
	}

	// 方式4 遍历数组  搭配range 关键
	var a [10]string
	a[0] = "Hello"
	for i := range a {
		fmt.Println("当前下标：", i)
	}
	for i, e := range a {
		fmt.Println("a[", i, "] = ", e)
	}

	// 遍历切片
	s := make([]string, 10) // 内建函数,对引用类型结构创建并返回初始化类型值
	s[0] = "Hello"
	for i := range s {
		fmt.Println("当前下标：", i)
	}
	for i, e := range s {
		fmt.Println("s[", i, "] = ", e)
	}

	m := make(map[string]string)
	m["b"] = "Hello, b"
	m["a"] = "Hello, a"
	m["c"] = "Hello, c"
	for i := range m {
		fmt.Println("当前key：", i)
	}
	for k, v := range m {
		fmt.Println("m[", k, "] = ", v)
	}

}

func loopBreak() {
	// -----------------------------------------------------------------------------------循环控制语句

	// 中断for循环
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("第", i, "次循环")
	}

	// 中断switch
	switch i := 1; i {
	case 1:
		fmt.Println("进入case 1")
		if i == 1 {
			break
		}
		fmt.Println("i等于1")
	case 2:
		fmt.Println("i等于2")
	default:
		fmt.Println("default case")
	}

	// 中断select
	select { // 常用监听多个 channel通道读写，作用于线程之间通信和同步
	case <-time.After(time.Second * 2):
		fmt.Println("过了2秒")
	case <-time.After(time.Second):
		fmt.Println("进过了1秒")
		if true {
			break
		}
		fmt.Println("break 之后")
	}

	// 不使用标记
	for i := 1; i <= 3; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			break
		}
	}

	// 使用标记  嵌套循环中，可以用 label 标出想 break 的循环。
outer: // 自定义 label 标签，标记外层想跳转循环位置
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			break outer // 跳出循环关键字,直接跳出被 outer标记的外层循环,而非只跳出内部循环
		}
	}

}

func loopContinue() {
	// ------------------------------------------------------------------------------------- continue 语句

	//	而且 continue 语句会执行 for 循环的 post 语句。
	//	在嵌套循环中，可以使用标号 label 标出想 continue 的循环。

	// 不使用标记
	for i := 1; i <= 2; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			if j >= 7 {
				continue
			}
			fmt.Println("不使用标记，内部循环，在continue之后执行")
		}
	}

	// 使用标记
outer:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			if j >= 7 {
				continue outer // 直接跳转到 outer位置 并继续向下执行
			}
			fmt.Println("不使用标记，内部循环，在continue之后执行")
		}
	}
}

func loopGoto() {
	// ----------------------------------------------------------------------------------- goto语句

	// goto 语句可无条件转移到指定 labal 标出的代码处。一般 goto 语句会配合条件语句使用，实现条件转移，构成循环，跳出循环的功能。
	//	一般不推荐使用 goto 语句，goto 语句会增加代码流程的混乱，不容易理解代码和调试程序。
	// 注意代码 label 声明之后，在代码中必须使用到，否则编译时会提示 label xxx defined and not used

	gotoPreset := false
preset:
	a := 5
process: // 自定义跳出 label
	if a > 0 {
		a--
		fmt.Println("当前a的值为：", a)
		goto process
	} else if a <= 0 {
		if !gotoPreset {
			gotoPreset = true
			goto preset
		} else {
			goto post
		}
	}
post:
	fmt.Println("main将结束，当前a的值为：", a)

}
