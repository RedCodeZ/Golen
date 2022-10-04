package main

import "fmt"

func main() {
	// Printf("格式化字符串"， 值)
	// %T 查看类型
	// %d 十进制数
	// %b 二进制数
	// %o 八进制数
	// %x 十六进制数
	// %c 字符
	// %s 字符串
	// %p 指针
	// %v 值
	// %f 浮点
 	var m1 = make(map[string]int, 1)
	m1["lixiang"] = 100
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1)
}
