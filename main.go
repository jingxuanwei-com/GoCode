package main

import "fmt"

func main() {
	//创建结构体
	var t1 Teacher
	fmt.Println(t1)
	t1.Name = "马泽"
	t1.Age = 45
	t1.School = "清华大学"
	fmt.Printf("%p\n", &t1)
	fmt.Println(t1)
	fmt.Println(t1.Name)

	var t2 = Teacher{"马碧", 44, "马龙大学"}
	fmt.Println(t2)

	var t3 = new(Teacher)
	(*t3).Name = "ssh"
	(*t3).Age = 45
	(*t3).School = "阿米诺斯"
	fmt.Printf("%p\n", &t3)
	fmt.Println(t3)
	fmt.Println(*t3)
}

// Teacher 结构体的定义
type Teacher struct {
	Name   string
	Age    int
	School string
}
