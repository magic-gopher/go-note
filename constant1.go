package main

import (
	"fmt"
)

func main() {
	/*
		常量：类似变量，也是一小块内存，用于存储数据，在程序运行的过程中，数据不可变。
		go语言的常量定义的关键字：const
		常数，固定不变的数值就是常数
	*/
	// 一个人的生日是固定不变的，可以定义为常量
	const BIRTHDAY = "1999.10.07"
	fmt.Println(BIRTHDAY)

	// 定义其他类型的常量
	const PI = 3.14
	fmt.Println(PI)

	// 常量定义之后，可以不使用
	const URL = "https://www.baidu.com:443"

	// 常量定义之后是不可以修改的
	// URL = "https://www.tianmao.com" // 尝试修改常量 // cannot assign to URL (untyped string constant "https://www.baidu.com:443")
	fmt.Println(URL)

	// 同时声明多个相同数据类型的常量
	const A, B, C = 100, 200, 300
	fmt.Println(A, B, C)

	// 同时声明多个不相同数据类型的常量
	const D, E, F = true, 6.66, "Golang"
	fmt.Println(D, E, F)

	// 定义一组常量
	const (
		a = 100
		b // 没有赋值，默认和上一行一样的数据类型和值
		c = "hello"
		d // 没有赋值，默认和上一行一样的数据类型和值
		e // 没有赋值，默认和上一行一样的数据类型和值
	)
	fmt.Printf("%T，%d\n", a, a)
	fmt.Printf("%T，%d\n", b, b)
	fmt.Printf("%T，%s\n", c, c)
	fmt.Printf("%T，%s\n", d, d)
	fmt.Printf("%T，%s\n", e, e)
}
