package main

import "fmt"

func main() {
	//Явное объявление с типом
	var name string = "Иван"
	var age int = 25

	//Тип выводится автоматичекски
	var city = "Москва" // string
	var height = 180.5  // Float64

	//Краткое объявление, работающее только внутри функции
	country := "China"
	weight := 75.3 // float64
	fmt.Printf(name, age, city, height, country, weight)

}
