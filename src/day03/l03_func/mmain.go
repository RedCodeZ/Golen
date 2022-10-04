package main

import "fmt"


func f1(){
	fmt.Println("沙河")
}

func f2(name string){
	fmt.Println("Hello", name)
}

//带参数和返回值
func f3(x int, y int) int {
	sum := x + y
	return sum
}

//参数类型简写
func f4(x, y int) int {
	sum := x + y
	return sum
}

//可变参数
func f5(x int, y ...int) int {
	fmt.Println(y)  //y沙河一个int类型的切片
	return 1
}

//命名返回值
// 带命名的函数，只能放在main函数外面。
func f6(x, y int) (sum int) { //在这里指定了sum的类型
	sum = x + y //如果使用命名的返回值，那么在函数中可以直接使用返回值变量
	return //return后面可以省略返回值变量sum
}

//go语言中支持多个返回值
func f7(x,y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {

	f1()
	f2("沙河")
	ret := f3(100, 200)
	fmt.Println(ret)
	ret2 := f5(1,2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(ret2)
}
