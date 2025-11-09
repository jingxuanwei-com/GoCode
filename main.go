package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println("自定义错误", err)
		//错误，强制程序中断
		panic(err)
	}
	fmt.Println("程序执行成功")
}

func test() error {
	a := 1
	b := 0
	if b == 0 {
		return errors.New("除数不能为0哦")
	} else {
		fmt.Println(a / b)
		return nil
	}
}
