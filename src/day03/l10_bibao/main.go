package main

import "fmt"

//闭包

func f1(f func())  {
	fmt.Println("This is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("This is f2")
	fmt.Println(x + y)
}
//现在要把f2传递到f1内，如何操作？
//定义一个函数队f2进行包装
//func f3 (x func(int, int), m, n int) func() {
//	tmp := func() {
//		x(m, n)
//	}
//	return tmp
//}

//在不改变f1的前提下，利用第三个函数把f2传递进了f1
//把原来需要传递两个int类型的参数包装成了一个不需要传递参数的函数
func f3(f func(int, int), x, y int)  func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}




func main() {
	//ret := f3(f2, 100, 200)
	////ret()
	//f1(ret)

	ret :=f3(f2, 100, 200)
	f1(ret)
}
