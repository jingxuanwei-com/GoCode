package main

import "fmt"

func main() {
	//打印函数
	fmt.Println(cal(10, 20))

	//直接使用函数
	calno(10, 50)

	//打印函数返回的两个值
	fmt.Println(cal2(10, 22))

	//提取两个值然后打印
	calz1, calz2 := cal2(10, 2)
	fmt.Println(calz1)
	fmt.Println(calz2)

	//只接受一个值，另一个_忽略
	_, calz3 := cal2(10, 2)
	fmt.Println(calz3)
}

// 新建函数cal 两个数相加
// 函数命名规范，
//首字母不能是数字
//首字母大写可以被本包文件和其他包文件使用
//首字母小写只能被本包使用

func cal(num1 int, num2 int) int {
	return num1 + num2
}

// 无返回值
func calno(num1 int, num2 int) {
	fmt.Println(num1 + num2)
}

func cal2(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 * num2
}
