package main

import "fmt"

func main() {
	var a = 2
	fmt.Printf("%d << 2 = 8 или 10 << 10 = %b\n", a, a<<2)

	//

	var a1 = 16
	var a2 = 3
	fmt.Printf("%d >> %d = 2 или %b >> %b = %b\n", a1, a2, a1, a2, a1>>a2)

	var a3 = 5
	var a4 = 2
	var a5 = 6

	var r1 = a3 | a4
	fmt.Printf("%d | %d = %d       %b | %b = %b\n", a3, a4, r1, a3, a4, r1)

	var r2 = a5 & a4
	fmt.Printf("%d & %d = %d       %b & %b = %b\n", a5, a4, r2, a5, a4, r2)

	var r3 = a3 ^ a4
	fmt.Printf("%d ^ %d = %d       %b ^ %b = %b\n", a3, a4, r3, a3, a4, r3)

	var r4 = a5 &^ a4
	fmt.Printf("%d &^ %d = %d       %b &^ %b = %b\n", a5, a4, r4, a5, a4, r4)

	var r6 = 0b1001_1101 &^ 0b1010_0101
	fmt.Printf("%d &^ %d = %d       %b &^ %b = %b\n",
		0b1001_1101, 0b1010_0101, r6,
		0b1001_1101, 0b1010_0101, r6) //0001_1000
	//0b1001_1101
	//0b1010_0101

	//0b0001_1000 &^
	//ob0111_1010 и-не

	// x y | z
	// 0 0 | 0
	// 0 1 | 0
	// 1 0 | 1    - СДНФ x & !y
	// 1 1 | 0	  - СКНФ (x | y) & (x | !y) & (!x | !y)

	var b1 = 5
	var c1 = 2
	var a7 = b1 | c1
	fmt.Printf("b1 : %b\n", b1)
	fmt.Printf("c1 : %b\n", c1)
	fmt.Printf("a7 : %b | %b = %b\n", b1, c1, a7)
}
