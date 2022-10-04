package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串只能使用双引号包括
	//单引号包括的是单独的字符
	//s1 := "hellp 沙河"
	//c1 := 'h'
	//c2 := '1'
	//c3 := '沙'

	//一个字节8bit，8个二进制
	//1个字符 A = 1个字节
	// 1个UTF8的字符，一个汉字是3个字符

	s3 := `D:\Go\src\code\day02`
	fmt.Println(s3)

	//字符串操作
	fmt.Println(len(s3))

	firstName := "louis"
	lastName := "Zhaung"

	//字符串拼接
	fullName := firstName + lastName
	fmt.Println(fullName)
	funcName2 := fmt.Sprintf("%s%s", firstName, lastName)
	fmt.Println(funcName2)

	//字符串分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	//包含,包含就返回true,否则就会返回false
	fmt.Println(strings.Contains(funcName2, "louis"))

	//前缀,判断一个字符串的前缀是否是以什么开头
	fmt.Println(strings.HasPrefix(funcName2, "louis"))

	//后缀，判断一个字符串的前缀是否是以什么结尾
	fmt.Println(strings.HasSuffix(funcName2, "louis"))

	//判断一个字符串出现的位置
	fmt.Println(strings.Index(funcName2, "u"))
	//判断一个字符串出现的位置,最后一次出现
	fmt.Println(strings.LastIndex(funcName2, "u"))

	//拼接,使用指定的符号连接
	fmt.Println(strings.Join(ret, "+"))
}
