package main

import "fmt"

// Go语言中的return不是原子操作，在底层分为两步来执行
// 第一步：返回值赋值
// 第二步：真正的RET返回
// 函数中的defer语句是在这两步之间执行的

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值=x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x  //返回值=y =x =5
}
func f4() (x int) {
	defer func(x int) {
		x++ //改变的是函数的副本
	}(x)
	return 5 //返回值=x=5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
