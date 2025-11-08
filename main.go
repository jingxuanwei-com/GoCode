package main

import "fmt"

func main() {
	//调用函数
	fmt.Println(cal(10, 20))
}

// 新建函数cal
func cal(num1 int, num2 int) int {
	return num1 + num2
}
