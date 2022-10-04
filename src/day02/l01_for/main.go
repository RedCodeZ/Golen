package main

import "fmt"

// for循环
func main() {

	//当 i=5,跳出本次for循环，不执行if
	//for i := 0; i < 10; i++ {
	//	if i == 5 {
	//		break
	//	}
	//	fmt.Println(i)
	//
	//}

	//跳出多层for循环的-常规方式
	//var flag = false
	//for i := 0; i < 10; i++ {
	//	for j := 'A'; j < 'Z'; j++ {
	//		if j == 'C' {
	//			flag = true
	//			break
	//		}
	//		fmt.Printf("%v-%c\n", i, j)
	//	}
	//	if flag {
	//      fmt.Println("over")
	//		break
	//	}
	//}

	//跳出多层for循环的-goto
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto jump
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}

	jump:
		fmt.Println("over")




}
