package main

import "fmt"

func main() {
	/*
		go语言的数据类型转换
		语法格式：Type(Value)

		注意事项：要是兼容的数据类型才能够进行转换

		常数：在有需要的时候，自动转型
		变量：需要手动转型
	*/
	var a int8
	a = 10

	var b int16
	b = int16(a) // 等同于java语言的自动类型转换
	fmt.Println(a, b)

	// 浮点类型
	f1 := 5.55
	c := int(f1) // 等同于java语言的强制类型转换
	fmt.Println(c)

	// 不兼容的数据类型是不能转换的
	// b1 := false
	// b2 := int8(b1) // cannot convert b1 (variable of type bool) to type int8
	// fmt.Println(b2)
}
