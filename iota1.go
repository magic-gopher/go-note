package main

import "fmt"

/*
iota：一种特殊的常量，可以被编译器修改的常量。
每定义一个const，iota初始值为0
每定义一个常量，就会自动累加1
一直到下一个const出现，清理
*/
func main() {
	const (
		a = iota // iota = 0
		b = iota // iota = 1
		c = iota // iota = 2
	)
	fmt.Println(a, b, c)

	const (
		d = iota    // iota = 0
		e = "hello" // iota = 1
		f = iota    // iota = 2
	)
	fmt.Println(d, e, f)
}
