package main

import "fmt"

// Вложенная структура (nested structure) - представляет собой поле внутри другой структуры, она имеет имя
// по которой можно обратиться
type person struct { //Вложенная структура
	name string
	age  int
}
type account struct {
	login       string
	password    string
	person_info person //Вложенная структура
}

// Встроенная структура (embedded structure) - представляет собой поле внутри другой структуры,
// но без имени, похожа на анонимное поле
type person1 struct { //Встроенная структура
	name string
	age  int
}
type account1 struct {
	login    string
	password string
	person1  //Встроенная структура
}

func main() {
	fmt.Println()

	tom := account{
		login:    "SomeLogin@gmail.com",
		password: "dsdasdasd",
		person_info: person{
			name: "Tom",
			age:  41,
		},
	}
	fmt.Println(tom)

	//Обращение к полям вложенной структуры
	fmt.Println("Name: ", tom.person_info.name)

	//Инициализация встроенной структуры

	var tom2 account1 = account1{
		"tom2@inbox.com",
		"sa2422",
		person1{"asdasda", 12},
	}
	fmt.Println(tom2)

	//Обращение к полям встроенной структуры
	fmt.Println("Age: ", tom2.person1.age)

	var node1 node = node{1, nil}
	var node2 node = node{2, nil}
	var node3 node = node{3, nil}

	node1.next = &node2
	node2.next = &node3

	printNodeValue(&node1)
}

// Хранение ссылки на структуру того же типа
type node struct {
	value int
	// next  node//Ошибка //Структрура не может иметь встроенно или вложенно
	// содержать в себе в точь-в-точь такой же тип
	// Вместо этого должно быть явно указана ссылка на ту же структуру в виде указателя
	next *node //ссылка на другую
}

// Рекурсивный вывод списка
func printNodeValue(node *node) {
	fmt.Println("Value: ", node.value, " Address: ", node, " Address next : ", node.next)
	if node.next != nil {
		printNodeValue(node.next)
	}
}
