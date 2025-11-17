package main

import "fmt"

func main() {
	var s1 Student
	var s2 Stu
	s1.Age = 10
	s2.Age = 20
	fmt.Println(s1)
	fmt.Println(s2)
	s1 = Student(s2)
	fmt.Println(s1)
}

type Student struct {
	Age int
}

type Stu struct {
	Age int
}
