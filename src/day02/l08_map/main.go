package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10) //初始化map，要估算好map的容量，避免在程序运行期间再动态扩容。
	m1["令狐冲"] = 18
	m1["岳不群"] = 32
	fmt.Println(m1)
	fmt.Println(m1["令狐冲"])
	fmt.Println(m1["娜扎"]) //如果不存在这个key，拿到对应值的类型的是0值，int是0, string是空字符串，bool是false
	value, ok := m1["娜扎"] //如果返回的值是bool,通常使用ok去接收它
	if !ok {
		fmt.Println("查无此key")
	}else {
		fmt.Println(value)
	}

	//map遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//遍历key
	for k := range m1 {
		fmt.Println(k)
	}

	//遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	//删除
    // 如果删除的key不存在，那么什么都不干。
	delete(m1, "岳不群")
	fmt.Println(m1)
	delete(m1, "任我行")
	fmt.Println(m1)
}

