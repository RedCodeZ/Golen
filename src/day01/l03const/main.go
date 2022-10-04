package main

//2.常量
//定义之后永远无法改变的量

const name = "louis"

const (
	//集中声明常量
	statusOk = 200
	notFont = 404
)

const (
	//集中声明常量
	//在声明常量时，如果后面的常量没有赋值，默认和上一个常量的值相同。
	n1 = 100
	n2
	n3
)

//iota
const (
	a1 = iota
	a2
	a3
)

func main()  {
	println(a1)
	println(a3)
}

