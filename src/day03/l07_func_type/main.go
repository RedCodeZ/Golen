package main

import "fmt"

func f1() {
	fmt.Println("沙河")
}

func f2() int {
	return 10
}

//函数可以作为参数的类型
func f3(x func() int) { //x是一个函数类型func(), func()int代表是一个int类型的函数
    ret := x()
	fmt.Println(ret)
}

func f4(x,y int) int {
	return x + y
}

func ff(a, b int) int  {
	return a + b
}

func f5(x func()int) func(int, int) int  { //函数的参数是一个int类型的参数，返回值是int类型的参数，这个返回值函数包含两个int类型的参数。
	return ff
}



func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	f3(f2)
	f3(b)
	f7 := f5(f2)
	fmt.Printf("%T\n", f7)

}
