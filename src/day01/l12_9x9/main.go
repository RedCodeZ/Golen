package main

import "fmt"

func main() {
	// 1 --------------------
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, j*i)
		}
		println()
	}
}
