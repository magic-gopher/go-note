package main

import "fmt"

func main() {
	/*
			关系运算符：关系运算符输出的结果只能是布尔类型
				>：表示大于
			    <：表示小于
				>=：表示大于等于
				<=：表示小于等于
				==：表示等于
		        !=：表示不等于
	*/
	a := 3
	b := 5
	c := 3
	res1 := a > b
	res2 := b > c
	// %t是格式化输出打印布尔类型的占位符
	fmt.Printf("%T,%t\n", res1, res1) // bool, false
	fmt.Printf("%T,%t\n", res2, res2) // bool, true
	res3 := a == b
	fmt.Printf("%T,%t\n", res3, res3) // bool, false

	fmt.Println(a != b, b != c, a != c)
}
