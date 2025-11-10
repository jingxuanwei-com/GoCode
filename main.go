package main

import "fmt"

func main() {
	var a [5]int
	//普通for循环
	for i := 0; i < len(a); i++ {
		fmt.Printf("输入第%d位学生的成绩：", i+1)
		fmt.Scanln(&a[i])
	}
	//for range循环
	for a, b := range a {
		fmt.Printf("第%d个学生成绩为%d\n", a, b)
	}
}
