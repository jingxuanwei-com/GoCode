package main

import (
	"fmt"
)

func main() {

	//创建方式2
	b := make(map[int]string)
	//增加
	b[100] = "Hello"
	b[200] = "World"

	//修改
	b[200] = "wd"
	fmt.Println(b)

	//删除
	delete(b, 200)
	fmt.Println(b)

	//查找
	value, flag := b[100]
	fmt.Println("查找的内容", value)
	fmt.Println("查找是否成功", flag)

	//获取长度
	fmt.Println("长度", len(b))

	//遍历
	for i, s := range b {
		fmt.Printf("第%d个是:%s\n", i, s)
	}

	//双重map
	c := make(map[string]map[int]string)
	c["a"] = make(map[int]string)
	c["a"][1555] = "A"
	c["a"][2666] = "B"
	c["a"][3777] = "C"
	c["a"][4888] = "D"
	fmt.Println(c)

	d := map[string]map[int]string{
		"a": {
			1555: "A",
			2666: "B",
		},
		"b": {
			1555: "A",
			2666: "B",
		},
	}
	for s, m := range d {
		fmt.Printf("学生班级为%v\n", s)
		for i, i2 := range m {
			fmt.Printf("学生学号为%v 姓名为%v\n", i, i2)
		}

	}
}
