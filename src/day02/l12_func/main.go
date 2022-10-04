package main

import "fmt"

//函数的定义
func sum(x int, y int)(ret int)  {
	ret = x + y
	return ret
}

//没有返回值
func f1(x int, y int)()  {
	fmt.Println(x + y)
}

//没有参数，没有返回值
func f2()  {
	fmt.Println("f2")
}


//没有参数，但是有返回值
func f3() int  {
	return 3
}

//参数可以命名，也可以不命名
//命名的返回值就相当于在函数中声明了一个变量
func f4(x int, y int)(ret int)  {
	ret = x + y
	return //使用命名返回值可以在return后面省略
}

//多个返回值
func f5() (int, string)  {
	return 1, "沙河"
}

//参数类型的简写
//当参数中连续多个参数的类型一致时，可以将非最后一个参数的类型省略
func f6(x, y int) int {
	return x + y
}
func f7(x, y int, m, n string, i, j bool) int  {
	return x + y
}

//可变长参数
//必须放在函数参数的最后
func f8(x string, y...int)  {
	//可以是多个同类型的参数
	//y的类型是切片
}


//Go语言中函数没有默认参数这个概念

func main() {

}
