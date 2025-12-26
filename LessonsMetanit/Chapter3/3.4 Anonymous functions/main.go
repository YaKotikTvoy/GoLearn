package main

import "fmt"

func main() {
	f := func(x, y int) int { return x + y }
	var t int = f(1, 2)
	fmt.Println(t)

	//Полное определение анонимной функции
	var some_func func(int, int) int = func(x, y int) (result int) {
		result = x + y
		fmt.Println(result)
		return
	}
	var a int = some_func(100, 2)
	fmt.Println(a)

	action(23, -9, func(x, y int) int { return x + y })
	action(23, -9, func(x, y int) int { return x * y })
}

func action(x, y int, operation func(int, int) int) {
	fmt.Println(operation(x, y))
}

// Анонимная функция как результат другой функции
func selectFn(n int) func(int, int) int {
	switch n {
	case 1:
		return func(x, y int) int { return x + y }
	case 2:
		return func(x, y int) int { return x - y }
	case 3:
		return func(x, y int) int { return x * y }
	default:
		return func(x, y int) int { return x / y }
	}
}
