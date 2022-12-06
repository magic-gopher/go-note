package main

import (
	"fmt"
)

func main() {
	/*
		字符串：一组byte的集合，可以理解为一组字符序列
		go语言定义字符串的方式：
			方式一："你好,gopher"
			方式二：`你好,gopher`
	*/

	// 定义一个字符串
	var s1 string
	s1 = "你好,gopher"
	fmt.Printf("%T,%s\n", s1, s1)

	// 'a'和"a"的区别
	s2 := 'a' // ''的只能存储一个字符，这里使用的''是int32数据类型
	s3 := "a"
	fmt.Printf("%T,%d\n", s2, s2)
	fmt.Printf("%T,%s\n", s3, s3)

	// 如何获取字符对应的int32类型的数值
	s4 := '你' + 0
	fmt.Printf("你，字符对应的数值是:%d\n", s4)
	// 将int32类型的数值转换为字符
	fmt.Println(string(s4))

	// 转移字符：字符串里面输出"需要通过\"来将其转义
	fmt.Println("\"hello world\"")
}
