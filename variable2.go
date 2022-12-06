package main

import "fmt"

/*
	变量的注意事项
	1、变量必须先定义才能使用
	2、变量的数据类型必须和赋值是一致的
	3、同一个作用域内变量的定义不能重复
	4、简短定义变量的方式，左边定义的变量至少有一个是新的变量
	5、全局变量不能使用简短定义的方式去定义
	6、go语言的变量定义了，必须使用，不如编译报错
	7、go语言的变量定义了，没有赋值，默认是有默认值的。
		7.1、数值型的默认值是：0
		7.2、浮点类型的默认值是：0.0
		7.3、布尔类型的默认值是：false
		7.4、字符串类型的默认值是：""
*/

// 声明一个全局变量a
var a bool = true

// b := 100 // 全局变量不支持使用简短定义的方式

func main() {
	var num = 100
	// %d是格式化打印输出整数类型的占位符
	// %p是格式化打印输出指针地址的占位符
	fmt.Printf("num的数值是%d,地址是%p\n", num, &num)

	num = 200 // 重新给变量赋值
	fmt.Printf("num的数值是%d,地址是%p\n", num, &num)

	// fmt.Println(name2) // undefined: name2

	var name string
	// name = 100 // cannot use 100 (untyped int constant) as string value in assignment
	name = "matrix"
	fmt.Println(name)

	// 使用简短定义的方式定义变量，左边的变量至少有一个变量是新的
	name, age, address := "Matrix", 18, "广东广州"
	fmt.Println(name, age, address)

	// 访问全局变量a
	fmt.Println(a)

	// 变量未赋值，会有默认值
	fmt.Println("-------------------")

	var m int
	fmt.Println(m) // 数值型的默认值是：0
	var n float64
	fmt.Println(n) // 浮点型的默认值是：0.0 -> 0
	var s string
	fmt.Println(s) // 字符串类型的默认值是：""
}
