package main

import "fmt"

func main() {
	var p Person
	p.Name = "John Doe"
	p.test()
	fmt.Println(p.Name)
	p.test1()
	fmt.Println(p.Name)
}

// Person 定义结构体
type Person struct {
	Name string
}

// 给Person结构体绑定方法test
func (s Person) test() {
	fmt.Println(s.Name)
	s.Name = "sssbbb"

}

// 给Person结构体绑定方法test,指针传递
func (s *Person) test1() {
	fmt.Println(s.Name)
	s.Name = "佳佳"
}
