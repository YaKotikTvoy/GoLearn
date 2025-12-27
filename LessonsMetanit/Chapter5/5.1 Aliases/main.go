package main

import "fmt"

//Определение псевдонима
// type name_alias тип_для_которого_определяем псевдоним

type mile uint
type kilometer uint
type binaryOp func(int, int) int

func main() {
	var distance mile = 5

	fmt.Println(distance)
	distance += 5
	fmt.Println(distance)

	//Для чего нужны псевдонимы, если они просто берут синонимизируют значения одного типа??
	// Как в данном случае mile и kilometer
	//Есть функция func distanceToEnemy(distance mile), принимающая именно mile

	distanceToEnemy(distance)

	//Если попытаться передать
	//var distance2 kilometer = 5
	//distanceToEnemy(distance2) // ошибка
	//Они позволяют явно разграничить смысл функций, повысить её описательность.

	var binaryop binaryOp = add
	action(1, 23823, binaryop)
}
func distanceToEnemy(distance mile) {
	fmt.Println("Расстояние до противника:")
	fmt.Println(distance, "миль")

}

//Также можно сократить название типов, если они слишком длинные
//	func action(n1, n2 int, op func(int, int) int) {
//		result := op(n1, n2)
//		fmt.Println(result)
//	}

func action(n1, n2 int, op binaryOp) {
	result := op(n1, n2)
	fmt.Println(result)
}
func add(x, y int) int {
	return x + y
}
