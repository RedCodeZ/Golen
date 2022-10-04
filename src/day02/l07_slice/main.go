package main

import "fmt"

//切片
func main()  {
	//切片的定义
	//var s1 [] int
	//var s2 [] string
	//fmt.Println(s1, s2)
	//fmt.Println(s1 == nil) //true 没有开辟内存空间
	//fmt.Println(s1 == nil)
	//
	////1.切片初始化
	//s1 = []int{1, 2, 3, 4, 5}
	//s2 = []string{"沙河", "张江", "平山村"}
	//fmt.Println(s1, s2)
	//
	////2.切片的长度和容量
	//fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	//fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	//3.由数组得到切片
	//a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	//s3 := a1[0:4] //基于数组切割，左包含右不包含，（左闭右开）
	//fmt.Println("s3", s3)
	//s4 := a1[1:6]
	//fmt.Println("s4:", s4)
	//s5 := a1[:4]
	//fmt.Println("s5:", s5)
	//s6 := a1[3:]
	//fmt.Println("s6:", s6)
	//s7 := a1[:]
	//fmt.Println("s7:", s7)

	//4.切片的容量
	//切片的额容量是指底层数组的容量
	//fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	////底层数组从切片开始到结束的容量
	//fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))
	//
	////5.切片再切片
	//s8 := s6[3:]
	//fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))


	//6.切片是一个引用类型
	//都指向了底层的一个数组
	//fmt.Println("s6:", s6)
	//a1[6] = 1300
	//fmt.Println("s6:", s6)
	//fmt.Println("s8:", s8)
	//
	////7.make函数创建切片
	//s10 := make([]int, 5, 10)
	//fmt.Printf("len(s10):%d cap(s10):%d\n", len(s10), cap(s10))
	//
	////8.切片赋值
	//s11 := []int{1, 3, 5, 7}
	//s12 := s11  //将s11直接赋值给s12，s11和s12共用一个底层数组
	//fmt.Println(s11, s12)
	//s12[1] = 300  //修改s11中第二个元素的值，那么s12中的第二个元素也会改变。因为s12和s11共用底层数组。
	//fmt.Println(s11, s12)

	//9.切片扩容 append()
	// 如果一个空的切片也是可以使用append去追加元素的，在追加元素的同时会把切片初始化。
	//s1 := []string{"北京", "上海", "广州"}
	////s1[3] = "深圳" //错误的写法，会导致编译错误，索引越界。
	////fmt.Println("s1")
	//fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//
	//// 调用append函数，必须用远来的切片变量接收返回值
	//// 必须用变量接收返回值
	//s1 = append(s1, "深圳") //append最佳元素，原来的底层数组放不下的时候，go底层就会把数组换一个。
	//fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//
	//s1 = append(s1, "重庆", "杭州", "西安")
	//fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//
	////把一个切片的内容追加到另一个切片
	//ss := []string{"北京", "武汉", "南京", "郑州", "厦门"}
	//s1 = append(s1, ss...) //...表示拆开追加
	//fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	//10.赋值切片
	//a1 := []int{1, 3, 5, 7, 9}
	//a2 := a1
	//var a3 = make([]int, 3, 3)
	//copy(a3, a1)
	//fmt.Println(a1, a2, a3)
	//a1[0] = 100
	//fmt.Println(a1, a2, a3)

	//11.删除切片元素
	// go语言中没有专用的方法来删除切片元素
	//将切片中1和3的元素删除
	//a1 := []int{1, 3, 5, 7, 9}
	//fmt.Println(a1[:1])
	//fmt.Println(a1[2:])
	//a1 = append(a1[:1], a1[2:]...)
	//fmt.Println(a1)
	//a1 = append(a1[:2], a1[3:]...)
	//fmt.Println(a1)

	//x1 := [...]int{1, 3, 5}
	//s1 := x1[:]
	//fmt.Println(s1)
	//s1 = append(s1[:1], s1[2:]...) //[1 5 5] 修改了底层数组前两个元素，第三个元素没有任何变化。
	//fmt.Println(s1) //[1 5]
	//fmt.Println(x1) //[1 5 5]

	//12.切片的练习
	var a = make([]int, 5, 10)
	fmt.Println(a)
	for i := 0; i < 10; i++{
		a = append(a, i)
	}
	fmt.Println(a)

	//var a = make([]string, 5, 10)
	//fmt.Println(a)
	//for i := 0; i < 10; i++{
	//	a = append(a, fmt.Sprintf("%v", i))
	//}
	//fmt.Println(a)
}
