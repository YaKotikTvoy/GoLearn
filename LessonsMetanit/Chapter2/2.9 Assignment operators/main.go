package main

import "fmt"

func main() {
	var a int = 10 // Простое присваивание
	b := 5

	a += b
	fmt.Println("a += b:", a) // 15

	a -= b
	fmt.Println("a -= b:", a) // 10

	a *= b
	fmt.Println("a *= b:", a) // 50

	a /= b
	fmt.Println("a /= b:", a) // 10

	a >>= b // a = a >> b
	fmt.Println("a >>= b:", a)

	a <<= b
	fmt.Println("a <<= b:", a)

	a |= b
	fmt.Println("a |= b:", a)

	a &= b
	fmt.Println("a &= b:", a)

	a ^= b
	fmt.Println("a ^= b:", a)

	a &^= b
	fmt.Println("a &^= b:", a)
}
