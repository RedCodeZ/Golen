package main

import "fmt"

// defer
// 一般用来函数结束前释放资源(关闭文件句柄，关闭数据库连接，关闭socket连接等操作)。
// 一个函数中可以有多个defer
// 多个defer语句按照先进后出的顺序延迟执行
func deferDemo() {
	fmt.Println("第一名")
	defer fmt.Println("第二名") //defer开头行的语句延迟到函数即将返回的时候再执行
	fmt.Println("第三名")
	defer fmt.Println("第四名")
}

func main() {
	deferDemo()
}