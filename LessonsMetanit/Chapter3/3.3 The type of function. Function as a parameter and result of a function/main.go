package main

import "fmt"

func main() {
	//переменная функция
	//(Это аналог делегатов в C#, только их нельзя обьявить как отдельный тип в Go)
	var f func(int, int) int = add
	var x = f(1, 2)
	fmt.Println(x)

	f1 := add
	fmt.Println(f1(3, 4))

	f1 = multiple
	fmt.Println(f1(3, 4))

	//display_delegate := display
	var display_delegate func(string) = display
	display_delegate("Some text")

	action(1, 2, add)
	action(15, 2, multiple)

	slice := []int{-2, 4, 5, -1, 7, -4, 23}
	fmt.Println(sum(slice, isEven))
	fmt.Println(sum(slice, isPositive))

	f = selectFn(1)
	fmt.Println(f(10, 2))

	f = selectFn(2)
	fmt.Println(f(10, 2))

	f = selectFn(3)
	fmt.Println(f(10, 2))
}

func add(x int, y int) (z int) {
	z = x + y
	return
}

func multiple(x int, y int) int {
	return x * y
}
func subtract(x, y int) int { return x - y }

func display(message string) {
	fmt.Println(message)
}

// Функция может передаваться в качестве параметров в другую ф-цию
func action(n1, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

func isEven(n int) bool {
	return n%2 == 0
}
func isPositive(n int) bool {
	return n > 0
}

func sum(numbers []int, criteria func(int) bool) int {
	result := 0
	for _, val := range numbers {
		if criteria(val) {
			result += val
		}
	}
	return result
}
func selectFn(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return subtract
	}
	return multiple
}
