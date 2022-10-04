package main

import "fmt"

//1.变量
//全局变量的声明
//同一个作用域中变量不能重复声明
//var name string
//var age int
//var status bool

//集中式声明变量
var (
	name string
	age int
	status bool
)

func main()  {

	//变量赋值
    name = "理想"
	age = 18
	status = true

	//简短变量声明
	//只能在函数内部使用
	//name := "louis"

	fmt.Printf("name: %s\n", name)
	fmt.Printf("age: %d\n", age)
	fmt.Println(status)

	//匿名变量

}