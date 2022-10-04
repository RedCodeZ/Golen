package main

import "fmt"
// 闭包
//一个函数引用了一个它外部作用域的变量，这就是闭包。
//1。函数可以作为返回值
//2。函数内部查找变量的顺序，现在自己内部找，如果找不到就往外层找。
func addr() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func addr2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	//addr:
	//ret := addr()
	//ret2 :=ret(200)
	//fmt.Println(ret2)

	//addr2:
	ret := addr2(100)
	ret2 := ret(200)
	fmt.Println(ret2)



}
