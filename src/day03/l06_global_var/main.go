package main

import "fmt"

// 变量的作用域
var x = 100 //定义一个全局变量

//定义一个函数
func f1() {
	//函数查找变量的顺序
	// 1.现在函数内部查找
	// 2.找不到就往函数外面查找，一直找到全局
	//函数内部定义的变量无法在函数外面使用
	fmt.Println(x)
}


func main() {
	f1()
}
