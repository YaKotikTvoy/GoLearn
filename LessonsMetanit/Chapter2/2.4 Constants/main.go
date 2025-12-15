package main

import "fmt"

//import "strconv"

func main() {
	const pi float64 = 3.1415

	const (
		pi2 float64 = 3.1415
		e           = 2.7182
	)

	const pi3, e2 = 3.1415, 2.7182
	const n = 5
	fmt.Println(pi)

	const (
		a = 1
		b
		c
		d = 3
		f
	)
	fmt.Println(a, b, c, d, f)

	//iota - идентификатор, увеличивающий с объявлением новой константы в инструкции определения множества
	//констант, но при объявлении новой инструкции объявления констант, она сбрасывается в ноль
	const ( //iota - сбрасывается в ноль
		C0 = iota // Здесь iota равна 0, но на новой строке объявления константын она увеличится на 1
		C1        // Увеличение iota на 1
		C2 = iota // Здесь iota уже 2
		C3        // хотя здесь константа становится 3...
		//C4 = "asjdlkasd " + iota. // Кажется, она просто увеличивает значение на один с объявлением новых констант
		C4
		C5
		C6 = iota // 6
	)

	const ( //iota - сбрасывается в ноль
		C7 = iota //0
	)

	fmt.Println("C0 : ", C0)
	fmt.Println("C1 : ", C1)
	fmt.Println("C2 : ", C2)
	fmt.Println("C3 : ", C3)
	fmt.Println("C4 : ", C4)
	fmt.Println("C5 : ", C5)
	fmt.Println("C6 : ", C6)
	fmt.Println("C7 : ", C7)
}
