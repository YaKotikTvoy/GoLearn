package main

import "fmt"

type person struct {
	name string
	age  int
}

// Данный метод получает копию структуры person
// И строка p.age = new_age, не изменит значение
// того экземпляра структры person, на которой был вызван метод, так как работа происходит не с ней
// а с копией
func (p person) update_age(new_age int) {
	p.age = new_age
}

// Для того чтобы изменить состояние структуры через метод,
// туда необходимо передать не копию а ссылку через указатель

func (p *person) update_age_ref(new_age int) {
	p.age = new_age
}

func main() {
	var tom person = person{name: "Tom", age: 24}
	fmt.Println("Before: ", tom.age)
	tom.update_age(25)
	fmt.Println("After: ", tom.age, "\n")

	fmt.Println("Before: ", tom.age)
	tom.update_age_ref(25)
	fmt.Println("After: ", tom.age)

	var tom2 *person = new(person)
	tom2.age = 23
	tom2.name = "Thomas"
	fmt.Println("Before: ", tom2.age)
	tom2.update_age_ref(34)
	fmt.Println("Before: ", tom2.age)
	tom2.update_age_ref(25)
	fmt.Println("After: ", tom2.age)

}
