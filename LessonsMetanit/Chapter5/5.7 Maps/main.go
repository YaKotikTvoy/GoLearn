package main

import (
	"fmt"
	"reflect"
)

func main() {
	// map - структура ведущая себя как словарь или хэш-таблица
	// хранит пары-ключ значение
	// map представляет тип map[K]V
	// где K - тип ключа
	// V - тип значения

	// общий вид объявления
	// var name_variable map[K]V = map[K]V{}

	// объявление с инициализацией
	// var name_variable map[K]V = map[K]V{
	//	ключ1 : значение1,
	//	ключ2 : значение2,
	//	ключ3 : значение3,
	//  ...
	//	ключN : значениеN,
	// }

	// есть одно ограничение для ключей, тип K - должен поддерживать операцию ==
	// это необходимо для нормальной работы хэш-таблицы

	// Значение карты поумолчанию равно nil, если ей ничего не присвоили
	var people map[string]int = map[string]int{
		"Tom":   12,
		"Sam":   24,
		"Alice": 23,
	}
	fmt.Println(people)

	// Обращение к элементам карты
	fmt.Println(people["Alice"])
	fmt.Println(people["Sam"])
	fmt.Println(people["Tom"])

	people["Tom"] = -238921893
	fmt.Println(people)

	// Проверка наличия ключа
	// Для проверки наличия элемента по определенному ключу можно применять
	// выражение if:
	// i, ok := map_name[key_name]
	// ключ существует, то ok будет true и значение найденного ключа
	// будет присвоено i, иначе ok = false, а значение i будет иметь
	// значение поумолчанию
	people = map[string]int{
		"Tom":       12,
		"Sam":       24,
		"Alice":     23,
		"Alexander": 120,
	}
	if val, ok := people["Alexander"]; ok {
		fmt.Println("val, ok := people[\"Alexander\"] ok =", ok, "val =", val)
	}

	if val, ok := people["Vera"]; ok {
		fmt.Println("val, ok := people[\"Vera\"] ok =", ok, "val =", val)
	} else {
		fmt.Println("val, ok := people[\"Vera\"] ok =", ok, "val =", val)
	}

	// Обход с помощью функции foreach
	for k, v := range people {
		fmt.Println(k, "-", v)
	}

	// Создание карты с помощью функции make
	people = make(map[string]int, 4)
	fmt.Println(len(people))
	fmt.Println(people)

	// Добавление нового элемента, происходит через индексатор
	people["Barsik"] = 29031
	people["Barsik1"] = 29031
	people["Barsik2"] = 29031
	fmt.Println(people)

	// Для удаления элемента применяют функцию
	// delete(map, key)
	delete(people, "Barsik")

	// Для сравнения двух map используется reflect.DeepEqual()

	// Две карты равны тогда, когда:
	// 1) когда у них одинаковы типы ключей и значений;
	// 2) одинаковое количество пар ключ-значение;
	// 3) соответствующие ключи равны по значением;
	// 4) значения по соответствующим ключам равны.
	people1 := map[string]int{"Tom": 1, "Sam": 2, "Bob": 12}
	people2 := map[string]int{"Tom": 1, "Sam": 2, "Bob": 3}
	people3 := map[string]int{"Tom": 1, "Sam": 2, "Bob": 12}
	people4 := map[int]string{1: "Tom", 2: "Sam", 12: "Bob"}
	fmt.Println("people1 == people2 :", reflect.DeepEqual(people1, people2)) // false
	fmt.Println("people2 == people3 :", reflect.DeepEqual(people2, people3)) // false
	fmt.Println("people1 == people3 :", reflect.DeepEqual(people1, people3)) // true
	fmt.Println("people4 == people3 :", reflect.DeepEqual(people4, people3)) // false
}