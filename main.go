package main

import "fmt"

func main() {
	num := 10
	//&返回变量指针位置
	fmt.Println(&num)
	fmt.Println(num)
	test(&num)
	fmt.Println(num)

}

// 参数类型为指针位置
func test(num1 *int) {
	//对地址相对应的变量做修改
	*num1 = 20
}
