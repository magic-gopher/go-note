package main

import "fmt"

func main() {
	/*
		赋值运算符
			=表示赋值操作
			+=表示加法操作再赋值
			-=表示减法操作再赋值
			*=
			/=
			%=
			<<=表示左移再赋值
			>>=表示右移再赋值
			^=
			&=
			|=
	*/

	// = 表示赋值操作
	var a int
	a = 10 // 将10数值赋值给a变量
	fmt.Println(a)

	a += 4         // a = a + 4
	fmt.Println(a) // a = 14

	a -= 3         // a = a - 3
	fmt.Println(a) // a = 11

	a *= 2         // a = a * 2
	fmt.Println(a) // a = 22

}
