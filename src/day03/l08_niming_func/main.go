package main

import "fmt"

//匿名函数


func main() {
	//函数内部是不能声明一个带名字的函数，只能定义匿名函数。
	//匿名函数，在函数内部定义函数时。
	//立即执行函数，只需要执行一次的函数。


	f1 := func (x, y int) { //把函数的结果赋值给一个变量，或者返回值。
		sum := x + y
		fmt.Println(sum)
	}

	//如果只是调用一次的函数，还可以简写成立即执行的函数
	func () {
		fmt.Println("Hello, Word")
	}() //这里加上括号，就表示立即执行。

	func (x, y int) {
		fmt.Println("Hello, Word")
		fmt.Println(x + y)
	}(100, 200) //带参数的函数可以在这里直接传参

	f1(10, 20)

}
