package main

import "fmt"

func main() {
	var intarr = [6]int{1, 2, 3, 4, 5, 6}
	slice := intarr[2:5]
	fmt.Println("intarr:", intarr)
	fmt.Println("slice:", slice)
	fmt.Println("slice len:", len(slice))
	fmt.Println("slice cap:", cap(slice))

	fmt.Printf("数组中下标为3位置的地址%p\n", &intarr[3])
	fmt.Printf("切片中下标为1位置的地址%p\n", &slice[1])

	//修改切片位置1的数值
	slice[1] = 16
	fmt.Println("intarr:", intarr)
	fmt.Println("slice:", slice)

}
