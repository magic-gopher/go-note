package main

import "fmt"

func main() {
	/*
					逻辑运算符：逻辑运算符操作的是布尔类型类型的数据
						&&：逻辑与，表示全部的条件都为真的，结果才是真的，只要有一个假的，结果就是假的。
				        ||：逻辑或，表示有只要有一个条件是真的，结果就是真的，只有全是假的，结果才是假的。
						!：逻辑非，表示取反，真的变成假的，假的变成真的。
					go语言也具备短路与和短路或运算
						例如：
							逻辑与：只要有一个是false，直接放回false
			                逻辑或：只要有一个条件是true，就直接放回true
					逻辑与口诀：一假全假，全真才真
		            逻辑或口诀：一真全真，全假才假

	*/

	f1 := true
	f2 := false
	f3 := true

	res1 := f1 && f2
	fmt.Printf("res1：%t\n", res1) // false

	res2 := f1 && f2 && f3
	fmt.Printf("res2：%t\n", res2) // false

	re3 := f1 || f2
	fmt.Printf("res2：%t\n", re3) // true

	res4 := f1 || f2 || f3
	fmt.Printf("res2：%t\n", res4) // true

}
