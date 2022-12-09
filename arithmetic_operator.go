package main

import "fmt"

func main() {
	/*
		算数运算符：+、-、*、/、%、++、--
	*/
	a := 10
	b := 3

	// 加法操作
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	// 减法操作
	sub := a - b
	fmt.Printf("%d - %d = %d\n", a, b, sub)

	// 乘法操作
	mul := a * b
	fmt.Printf("%d * %d = %d\n", a, b, mul)

	// 除法操作
	div := a / b
	fmt.Printf("%d / %d = %d\n", a, b, div)

	// 取余操作
	mod := a % b
	fmt.Printf("%d %% %d = %d\n", a, b, mod) // %%在格式化打印中表示一个%

	c := 3
	c++ // ++ 操作在golang里面不支持复杂的算数运算
	fmt.Println(c)
}
