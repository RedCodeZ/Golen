package main

import "fmt"

func main()  {
	var a1 = [3] string {"北京", "上海", "广州"}
	fmt.Println(a1)
	a1[1] = "天津"
	fmt.Println(a1)
	a2 := a1[:1]
	fmt.Println(a2[:1])
	fmt.Printf("type: %T\n", a2)
	fmt.Printf("type: %T\n", a1)
}
