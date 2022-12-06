package main

import "fmt"

/*
	变量：其实就是一小块内存，用于存储数据，在程序运行的过程中，数据是可变的。
	go语言中变量定义的关键字：var
	go语言中定义变量的三种方式：
		第一种方式：先定义后赋值
			var 变量名称 数据类型
			变量名称 = 数值
		第二种方式：类型推断
			var 变量名称 = 数值
		第三种方式：简短定义(这种方式相当于是重新定义了一个新的变量，没有赋值的操作，使用这种方式定义变量，必须有一个变量是新的)
			变量名称 := 数值
*/

func main() {
	// 第一种方式
	var num1 int
	num1 = 100
	fmt.Printf("num1的数值是:%d\n", num1)

	// 第一种方式写在一行
	var num2 int = 1243
	fmt.Printf("num2的数值是:%d\n", num2)

	// 第二种方式
	var name = "matrix"
	fmt.Printf("数据类型是%T,数值是%s\n", name, name)

	// 第三种方式
	sum := 1000
	fmt.Println(sum)

	// 同时声明多个变量
	var a, b, c int // 定义多个相同数据类型的变量
	a = 10
	b = 20
	c = 30
	fmt.Println(a, b, c)

	// 同时定义多个不相同的数据类型的变量
	var m, n, f = "a", 100, false
	fmt.Println(m, n, f)

	// 定义一组变量，变量的数据类型可以不同
	var (
		StudentName    string = "matrix"
		StudentAge     int    = 18
		StudentAddress string = "广东广州"
	)
	fmt.Println(StudentName, StudentAge, StudentAddress)
}
