package main

import "fmt"

func main() {
	//	指定变量类型并且赋值
	var num1 int = 18
	fmt.Println(num1)

	//	指定变量类型，但不赋值，使用默认值
	var num2 int
	fmt.Println(num2)

	//	自动类型推断
	var num3 = "tom"
	fmt.Println(num3)

	//	省略var =写为:=
	num4 := "男"
	fmt.Println(num4)
}
