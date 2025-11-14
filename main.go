package main

import (
	"fmt"
)

func main() {
	//直接创建切片 切片类型 切片长度 切片容量
	slice := make([]int, 10, 20)
	fmt.Println(slice)
	fmt.Println("切片长度:", len(slice))
	fmt.Println("切片容量:", cap(slice))

	slice[0] = 66
	slice[1] = 88
	fmt.Println("slice", slice)

	//定义数组
	var intarr = [6]int{1, 2, 3, 4, 5, 6}
	//定义切片
	var slice2 = intarr[1:2]
	slice2[0] = 66
	fmt.Println("intarr", intarr)
	fmt.Println("slice2", slice2)

	i := 10
	//底层追加元素的时候进行扩容，老数组扩容为新数组
	slice3 := append(slice2, i)
	fmt.Println("slice3", slice3)

	//给自己追加
	//底层是新数组，不能直接维护，需要通过切片间接维护操作
	slice3 = append(slice3, 11)
	fmt.Println("slice3", slice3)
	//...代表追加是是一个切片
	slice3 = append(slice3, slice2...)
	fmt.Println("slice3", slice3)

	//拷贝
	copy(slice3, slice)
	fmt.Println("slice3", slice3)

}
