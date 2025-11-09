package main

import "fmt"

func main() {
	test()
	fmt.Println("程序执行成功")
}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("test发生错误", err)
		}
	}()
	a := 1
	b := 0
	fmt.Println(a / b)
}
