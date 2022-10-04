package main

import "fmt"

func main() {
	c1 := "H"
	c2 := 'H'
	c3 := "沙"
	c4 := '沙'
	c5 := byte('H')

	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("c5:%T\n", c5)

	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))


	//
}
