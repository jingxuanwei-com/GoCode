package main

import "fmt"

func main() {
	score := 87
	switch score / 10 {
	case 10:
		fmt.Println("You are 10%")
	case 9:
		fmt.Println("You are 9%")
	case 8:
		fmt.Println("You are 8%")
	case 7:
		fmt.Println("You are 7%")
	case 6:
		fmt.Println("You are 6%")
	case 5:
		fmt.Println("You are 5%")
	case 4:
		fmt.Println("You are 4%")
	case 3:
		fmt.Println("You are 3%")
	case 2:
		fmt.Println("You are 2%")
	case 1:
		fmt.Println("You are 1%")

	default:
		fmt.Println("You are 0%")
	}
}
