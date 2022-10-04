package main

import "fmt"

//浮点数
func main() {
	f1 := 1.234567
	fmt.Printf("%f\n", f1) //默认Go语言中的小数都是float64类型的

	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)
	//f1 = f2 //float32类型的值不能直接赋值给float64的类型
}
