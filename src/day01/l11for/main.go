package main

import "fmt"

//for循环
func main() {
	
	//基本格式
	//s := "hello"
	//for i, v := range s {
	//	fmt.Printf("%d %c\n", i, v)
	//}

	//哑元变量，不想用到额都直接扔给它
	s := "hello"
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}
}
