package main

import "fmt"

func main() {
	/*
		go语言的数据类型
		1、基本类型
			布尔类型
			整型
			浮点型
			复数
			字符
			字符串
		2、复合类型（派生类型）
	*/

	// 布尔类型
	var b1 bool
	b1 = true
	fmt.Printf("%T,%t\n", b1, b1) // %t是格式刷打印布尔类型的值的占位符
	b2 := false
	fmt.Printf("%T,%t\n", b2, b2)

	// 整数（u开头的都是无符号位的，都是从0开始）
	var i1 int8
	i1 = 100
	fmt.Println(i1)
	var i2 uint8
	i2 = 200
	fmt.Println(i2)

	// 这里的int是根据操作系统的位数决定的，我的是64位的所以，int就是64位的，但是int数据类型不等于int64
	var i3 int
	i3 = 1000
	fmt.Println(i3)

	var i4 int64
	// 从golang的语法角度说明int在64位的系统时和int64不是一种数据类型（这里需要转换）
	// i4 = i3 // cannot use i3 (variable of type int) as type int64 in assignment
	fmt.Println(i4)

	// unit8和byte类型是兼容的
	var i5 uint8
	i5 = 100
	var i6 byte
	i6 = i5
	fmt.Println(i5, i6)

	// 浮点类型
	var f1 float32
	f1 = 3.14
	fmt.Printf("%T,%0.2f\n", f1, f1) // %f是格式化打印浮点类型的占位符，0.2表示只输出小数点后两位
	var f2 float64
	f2 = 6.123456
	fmt.Printf("%T,%0.2f\n", f2, f2)
	f3 := 10.24
	fmt.Printf("%T,%0.2f\n", f3, f3) // 浮点型默认是float64类型
}
