package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println(numbers)

	var numbers2 [6]string
	fmt.Println(numbers2)

	//Можно определить меньше элементов, но не больше
	var someNumber = [5]int{1, 2, 3}
	fmt.Println(someNumber)

	//Краткое определение массива
	numbers3 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers3)

	//Можно не прописывать явно количество объектов, это можно сделать через троеточие,
	//компилятор сам определит, сколько памяти точно нужно выделить в массиве
	numbers4 := [...]byte{232, 23, 1, 232}
	var numbers5 = [...]byte{232, 23, 1, 232}
	//Так нельзя, определение с не точным определением справа, недопускает точное определение слева
	//var numbers6 []byte = [...]byte{232, 23, 1, 232}
	numbers6 := [...]string{"Барсик", "Бусик", "Mursik"}

	fmt.Println(numbers4, numbers5, numbers6)

	//Массивы в go нельзя перепресваивать
	numbers7 := [...]int32{23232, 232323, -2323}
	numbers8 := [...]int32{-2323, 2574323, -0023}
	//numbers7 = numbers8 // ошибка
	fmt.Println(numbers7, numbers8)

	//Индексы
	numbers9 := [8]uint{1, 2, 2, 3, 4, 5, 6, 2323}
	numbers9[2] = numbers9[7]
	fmt.Println(numbers9[2])
	fmt.Println(numbers9[7])

	// Длина массива является частью его типа, [3]int != [4]int

	var numbers10 [3]int = [3]int{1, 2, 3}
	var numbers11 [4]int = [4]int{1, 2, 3, 4}
	// numbers10 = numbers11 // Ошибка
	fmt.Println(numbers10, numbers11)

	// Индексы - ключи, и им можно явно задавать значения и необязательно по порядку
	colors := [3]string{1: "blue", 0: "red", 2: "green"}
	fmt.Println(colors)
	fmt.Println(colors[2]) // green

	// Многомерные массивы
	// Имя_массива := [размерность_1][размерность_2]..тип_массива{значения_массива}
	numbers12 := [3][2]int{{1, 2}, {4, 5}, {7, 8}}
	fmt.Println(numbers12)

	// Как и в случае с одномерными массивами нет необходимости инициализировать
	// все элементы многомерного массива, остальные элементы будут проинициализированы
	// по умолчанию
	numbers13 := [3][2]int{
		{1, 2},
		{5},
	}
	fmt.Println(numbers13)

	var number14 [3][2]int
	number14[0] = [2]int{1, 2}
	number14[1] = [2]int{4, 5}
	number14[2] = [2]int{7, 8}

	fmt.Println(number14)

	var number15 [3][2]int
	number15[0][0] = 1
	number15[0][1] = 2
	number15[1][0] = 4
	number15[1][1] = 5
	number15[2][0] = 7
	number15[2][1] = 8

	fmt.Println(number15)

	//для того чтобы узнать длину массива используют функцию len()
	//Так можно объявить многомерный массив без задания размера
	numbers16 := [...][]int{{1, 2}, {2, 2}, {3, 2}, {1, 1}}
	fmt.Println(len(numbers16))

	//Сравнение массивов
	//Два массива равны друг другу, если они имеют
	// одинаковый тип
	// размерность
	// одинаковое количество элементов
	// каждый элемент соответствующего адреса равен каждому элементу другого массива этого же адреса

	nums1 := [4]int{3, 4, 5, 6}
	nums2 := [4]int{3, 4, 5}
	fmt.Printf("nums1 == nums2 : %t\n", nums1 == nums2) // false

	nums3 := [3][2]int{{2}, {5}}
	nums4 := [3][2]int{{2, 1}, {5}}
	fmt.Printf("nums3 == nums4 : %t\n", nums3 == nums4) // false

	nums5 := [4]int{3, 4, 5, 0}
	fmt.Printf("nums5 == nums2 : %t\n", nums5 == nums2) // true

	// Копирование массива
	// Массивы в Go ведут себя как значения и при приравнивании одного массива другому
	// изменение значений в одном массиве никак не отразится на другом

	nums6 := [3][2]int{{1, 2}, {4, 5}, {7, 8}}
	nums7 := nums6
	nums7[2][0] = 28392983
	fmt.Println("nums7[2][0] :", nums7[2][0])
	fmt.Println("nums6[2][0] :", nums6[2][0])
}
