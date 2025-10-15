package main

import (
	"fmt"

	_ "github.com/fan-chao-sys/GoBase/pkg1"
)

const mainName string = "main"

var mainVar string = getMainVar()

func init() {
	fmt.Println("main init method invoked")
}

func main() {
	fmt.Println("main method invoked!")
	fmt.Println("mainVar:", mainVar)
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked!")
	return mainName
}
