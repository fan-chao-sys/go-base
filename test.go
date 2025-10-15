// 包声明
package main

// 引入包声明, 单个不带括号，多个括号逗号隔开
import (
	"fmt"
)

// PrintInConsole 函数声明
// PrintInConsole 将字符串打印到控制台
func PrintInConsole(s string) {
	fmt.Println(s)
}

// Str 全局变量声明
// 全局变量示例（保持导出以便其他包访问）
var Str string = "ASDFBF"

// 主函数启动类
// init 保持空实现
func init() {

}

// 系统初始化函数
func init() {

}
