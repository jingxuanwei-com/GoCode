package main

import "fmt"

func main() {
	fmt.Println(test(10, 20, 50, 30))
}

// 可变参数... 可以传入多个参数、
// len()获取数组长度
func test(args ...int) int {
	num1 := 0
	for i := 0; i < len(args); i++ {
		num1 = num1 + args[i]
	}
	return num1
}
