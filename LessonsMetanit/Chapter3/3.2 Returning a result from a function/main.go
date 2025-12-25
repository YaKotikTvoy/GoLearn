package main

import "fmt"

//func имя_функции (список_параметров) (типы_возвращаемого_значения) {
//    выполняемые_операторы
//    return возвращаемое_значение
//}

func main() {
	var a = add2(1, 2)
	fmt.Println(a)

	age, name := add3(4, 5, "Tom", "Simpson")
	fmt.Println(age)  // 9
	fmt.Println(name) // Tom Simpson

	age2, name2 := add4(4, 5, "Tom", "Simpson")
	fmt.Println(age2)  // 9
	fmt.Println(name2) // Tom Simpson
}

func add(x, y int) int {
	return x + y
}

// Функция может возвращать именнованный возвращаемый результат
func add2(x, y int) (z int) {
	z = x + y
	return
}

// Функция может возвращать несколько значений
func add3(x, y int, firstName, lastName string) (int, string) {
	var z = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

// Именнованные параметры
func add4(x, y int, firstName, lastName string) (z int, fullName string) {
	z = x + y
	fullName = firstName + " " + lastName
	return
}
