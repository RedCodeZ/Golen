package main

import "fmt"

func main() {
	//记得要初始化
	//元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "沙河"
	fmt.Println(s1)

	//值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30, 40} //初始化
	fmt.Println(m1)

	//var mapSlice = make([]map[string]string,3)
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}
	//fmt.Println("after init")
	//// 对切片中的map元素进行初始化
	//mapSlice[0] = make(map[string]string, 10)
	//mapSlice[0]["name"] = "小王子"
	//mapSlice[0]["password"] = "123456"
	//mapSlice[0]["address"] = "沙河"
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}
}
