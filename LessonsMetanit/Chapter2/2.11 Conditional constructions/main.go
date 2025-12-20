package main

import "fmt"

func main() {
	a := 6
	b := 7
	if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a больше b")
	}

	if a < b {
		fmt.Println("a less than b")
	} else if a == b && a > b {
		fmt.Println("a greater or equal than b")
	} else if b == a && b > a {
		fmt.Println("b greater or equal than a")
	} else {
		fmt.Println("a == b")
	}

	a = 6
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	case 6:
		fmt.Println("a = 6")
		fallthrough
	case 5, 4, 3:
		fmt.Println("a = 5 || a = 4 || a = 3")
		fallthrough
	case 2:
		fmt.Println("a = 2")
		fallthrough // провал через, работа переходит на блок case ниже
	default:
		fmt.Println("Значение переменной а неопределенно")
	}
}
