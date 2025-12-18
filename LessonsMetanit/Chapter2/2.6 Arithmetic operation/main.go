package main

import "fmt"

func main() {
	var a = 4
	var b = 6
	var c = a + b
	fmt.Println(c) // 10

	c = a - b
	fmt.Println(c) // -2

	c = a * b
	fmt.Println(c) // 24

	//Деление целых чисел возвращает целое число, даже если оно будет присвоено вещественному типу
	c = a / b
	fmt.Println(c)     // 2
	fmt.Println(a / b) // 2

	var k float32 = 10
	var l float32 = 4
	var m float32 = k / l
	fmt.Println(m) // 2.5

	var divide float32 = 10.0 / 4.0
	fmt.Println(divide)

	//Остаток от деления
	var reclame int32 = 35 % 3
	fmt.Println(reclame)

	//Декременты и инкременты, постфиксные и префексные
	var increment = 2
	increment++
	//var increment2 = 2 + increment++ // Like this wrong
	fmt.Println(increment)
	increment--
	fmt.Println(increment)
}
