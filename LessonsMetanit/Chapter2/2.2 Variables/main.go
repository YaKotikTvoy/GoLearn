package main

import "fmt"

func main() {
	var hello string
	var a, b, c string
	hello = "Hello world"
	fmt.Println(hello, a, b, c)

	var hello2 string = "Intialize variable"
	fmt.Println(hello2)

	//Множественное определение с указанием типов и значений
	var (
		name string = "Tom"
		age  int    = 27
	)
	fmt.Println(name, age)

	// Множественное определение с автоматическим определением типа через
	// инициализацию
	var (
		name2 = "Что-то"
		age2  = 23
	)
	fmt.Println(name2, age2)

	//Краткое определение переменной, так можно только в функции
	name3 := "Barsik"
	fmt.Println(name3)
}
