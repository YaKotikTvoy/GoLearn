package main

import "fmt"

// Методы в Go представляют из себя также функции, но они связаны с определенным типом,
// они вызываются на определенном экземпляре структуры

/*
Общий вид объявления

func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возращаемых_результатов) {

}
*/

// допустим это именнованный тип, представляющий срез из строк
type library []string

// Для вывода всех элементов из среза
func (l library) print() {
	for _, v := range l {
		fmt.Println(v)
	}
}

func main() {
	var lib_ = [...]string{"Book1", "Book2", "Book3"}
	for _, v := range lib_ {
		fmt.Println(v)
	}

	var lib library = library{"Book1", "Book2", "Book3"}
	lib.print()

	var tom person = person{"Tom", 24}
	tom.eat("soup")
	tom.print()
}

// Struct methods
func (p person) print() {
	fmt.Println("Name : ", p.name)
	fmt.Println("Name : ", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "eat", meal)
}

type person struct {
	name string
	age  int
}
