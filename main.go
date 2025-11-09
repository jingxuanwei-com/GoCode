package main

import "fmt"

func main() {
	//定义一个函数类型的变量
	a := test
	fmt.Printf("a 的变量类型是 %T\n", a)
	a(1)
	b(2, 3, a)

	// 自定义数据类型，别名
	type myInt int
	var sb1 myInt = 42
	fmt.Println(sb1)

}

func test(a int) {
	fmt.Println(a)
}

// 可传入函数类型变量
func b(a int, b int, c func(int)) {
	c(a + b)
}
