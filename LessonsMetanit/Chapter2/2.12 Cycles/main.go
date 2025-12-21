package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i * i)
	}
	var i = 1
	for ; i < 10; i++ {
		fmt.Println(i * i)
	}
	for i = 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	str := "Hello"
	//foreach in Go
	for index, value := range str {
		fmt.Printf("Index: %d Value: %c\n", index, value)
	}

	// Если не планируем использовать значения или индексы элементов,
	// то можно вместо них указать прочерк
	for _, value := range str {
		fmt.Printf("Value: %c\n", value)
	}
	fmt.Printf("%c\n", 10)

	var users = [3]string{"Tom", "Alice", "Kate"}
	for index, value := range [3]string{"Tom", "Alice", "Kate"} {
		fmt.Println(index, value)
	}

	//Для перебора можно использовать и стандартную версию цикла for
	for i = 0; i < len(users); i++ {
		fmt.Printf("index : %d Value : %s\n", i, users[i])
	}

	//continue
	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum = 0

	for _, value := range numbers {
		if value < 0 {
			continue //
		}
		sum += value
	}
	fmt.Println(sum)

	//break

	numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum = 0

	for _, value := range numbers {
		if sum > 4 {
			break
		}
		sum += value
	}
	fmt.Println("Sum:", sum) // Sum : 10

	//Оператор break может осуществить выход по метке, как goto
	//метка для выхода к внешнему циклу
Outerloop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {

			fmt.Printf("i = %d, j = %d\n", i, j)
			//if i == 1 && j == 2 {
			//	fmt.Println("Выход из внешнего цикла...")
			//	goto Outerloop2 // Выходим из внешнего цикла
			//}

			if i == 2 && j == 2 {
				fmt.Println("Выход из внешнего цикла...")
				break Outerloop // Выходим из внешнего цикла
			}
		}
	}
	//Outerloop2:
	fmt.Println("Цикл завершен...")
}
