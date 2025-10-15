// 包声明
package GoBase

// 引入包声明, 单个不带括号，多个括号逗号隔开
import (
	"fmt"
)

// 函数声明
func printInConsole(s string) {
	fmt.Println(s)
}

// 全局变量声明
var str string = "ASDFBF"

// 主函数启动类
func main() {
	fmt.Println(str)
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
}

// 系统初始化函数
func init() {

}
