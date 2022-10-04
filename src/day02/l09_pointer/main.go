package main

import "fmt"

func main() {
	//1. &:取内存地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)  //int类型的指针
	//2. *:根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

}
