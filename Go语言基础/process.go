package Go语言基础

import "fmt"

func processIf() {
	// ------------------------------------------------------------------------------  if 语句
	// if 语句由一个或多个布尔表达式组成，且布尔表达式可以不加括号

	// 嵌套格式:
	// if <expression1> {
	//    if <expression2> {
	//        <do sth1>
	//    } else {
	//        <do sth2>
	//    }
	//} else if <expression3> {
	//    <do sth3>
	//} else {
	//    <do sth4>
	//}

	// if/else 语句可在布尔表达式前额外增加声明赋值语句，来声明作用域仅在当前 if 作用域内的变量：
	var a int = 10
	if b := 1; a > 10 {
		b = 2
		fmt.Println("a > 10")
	} else if c := 3; b > 1 {
		b = 3
		fmt.Println("b > 1")
	} else {
		fmt.Println("其他")
		if c == 3 {
			fmt.Println("c == 3")
		}
		fmt.Println(b)
		fmt.Println(c)
	}

}

func processSwitch() {
	// ----------------------------------------------------------------------------------  switch 语句
	// 默认情况下 case 分支自带 break 效果，无需在每个 case 中声明 break，中断匹配。

	a := "test string"
	// 1. 基本用法
	switch a {
	case "test":
		fmt.Println("a = ", a)
	case "s":
		fmt.Println("a = ", a)
	case "t", "test string": // 可以匹配多个值，只要一个满足条件即可
		fmt.Println("catch in a test, a = ", a)
	case "n":
		fmt.Println("a = not")
	default:
		fmt.Println("default case")
	}

	// 变量b仅在当前switch代码块内有效
	switch b := 5; b {
	case 1:
		fmt.Println("b = 1")
	case 2:
		fmt.Println("b = 2")
	case 3, 4:
		fmt.Println("b = 3 or 4")
	case 5:
		fmt.Println("b = 5")
	default:
		fmt.Println("b = ", b)
	}

	// 不指定判断变量，直接在case中添加判定条件
	b := 5
	switch {
	case a == "t":
		fmt.Println("a = t")
	case b == 3:
		fmt.Println("b = 5")
	case b == 5, a == "test string":
		fmt.Println("a = test string; or b = 5")
	default:
		fmt.Println("default case")
	}

	var d interface{}
	d = 1
	// 对于 d的类型进行判断方式
	switch t := d.(type) {
	case byte:
		fmt.Println("d is byte type, ", t)
	case *byte:
		fmt.Println("d is byte point type, ", t)
	case *int:
		fmt.Println("d is int type, ", t)
	case *string:
		fmt.Println("d is string type, ", t)
	case *CustomType:
		fmt.Println("d is CustomType pointer type, ", t)
	case CustomType:
		fmt.Println("d is CustomType type, ", t)
	case int:
		fmt.Println("d is int type, ", t)
	default:
		fmt.Println("d is unknown type, ", t)
	}
}

type CustomType struct{}
