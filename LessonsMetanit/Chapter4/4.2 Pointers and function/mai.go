package main

import "fmt"

func main() {
	d := 3
	fmt.Println("d before : ", d)
	changeValue(d)
	fmt.Println("d after : ", d)

	fmt.Println("d before : ", d)
	RealChangeValue(&d)
	fmt.Println("d after : ", d)

	p := createPointer(7)
	fmt.Println(*p)

	p2 := createPointer(72)
	fmt.Println(*p2)

	p3 := createPointer(72)
	fmt.Println(*p3)
}

// Функция принимает копию переменной, но не её
func changeValue(x int) {
	x *= x
}

// При подаче ссылки, а не значения, изменяется непосредственно переменная, а не её копия
// с другой стороны это даёт прирост скорости, так как происходит обращение по адресу
// а не передача копии, которая может быть огромной по памяти
func RealChangeValue(x *int) {
	*x *= *x
}

func createPointer(i int) (p *int) {
	*p = i
	return
}
