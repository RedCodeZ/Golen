package main

import "fmt"

//占位符
func main() {
	var n = 100
	var s = "hello 沙河"
	//查看类型
	fmt.Printf("%T\n",n) //查看类型
	fmt.Printf("%v\n",n) //查看变量的值
	fmt.Printf("%b\n",n) //转换为二进制
	fmt.Printf("%d\n",n) //转换为十进制
	fmt.Printf("%o\n",n) //转换为八进制
	fmt.Printf("%x\n",n) //转换为十六进制
	fmt.Printf("%s\n",s) //查看字符串
	fmt.Printf("%v\n",s) //查看变量的值
	fmt.Printf("%#v\n",s)
}
