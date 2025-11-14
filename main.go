package main

import "fmt"

func main() {
	//创建方式1
	//定义map变量
	var a map[int]string
	//必须通过make初始化才会分配空间
	a = make(map[int]string)

	a[100] = "Hello"
	a[200] = "World"
	fmt.Println(a)

	//创建方式2
	b := make(map[int]string)
	b[100] = "Hello"
	b[200] = "World"
	fmt.Println(b)

	//方式3
	c := map[int]string{
		100: "Hello",
		200: "World",
	}
	fmt.Println(c)
}
