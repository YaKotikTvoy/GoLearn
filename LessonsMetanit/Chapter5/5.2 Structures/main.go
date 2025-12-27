package main

import (
	"fmt"
)

/*
Общее определение структуры

	type имя_структуры struct{
		поле1 тип_поля1
	    поле2 тип_поля2
	    ...............
	    полеN тип_поляN
	}
*/

type person struct {
	name string
	age  int
}

// Если структура имеет поле одного типа, то их можно совместить
type person2 struct {
	name, company string
	age           int
}

//Anonymous struct

var tom struct {
	name string
	age  int
}

var tom6 = struct {
	name string
	age  int
}{ //Инициализация структуры
	name: "Tom",
	age:  9,
}

// Анонимные поля структуры
// необязательно объявлять имена полям, если типы полей в структуры имеются только один раз
// в таком случае имя поля совпадёт с типом поля

type anonymousFieldInStruct struct {
	string
	bool
	int8
	int16
	int
	rune
}
type anon anonymousFieldInStruct

//Структуру можно определять как и вне функций, на уровне пакета, так и внутри функций

func main() {

	type person3 struct {
		name, company string
		age           int
	}

	//Инициализатор структуры - type{"Поля", " по ", " инициализируются по соответствующим местам", 23}
	var tom person3 = person3{"Tom", "BarsikSoft", 23}
	fmt.Println(tom)

	//Можно явно указать какие значения передаются свойствам структуры

	alice := person3{name: "Alice", age: 24, company: "Google"}
	fmt.Println(alice)

	//Можно даже не указывать никаких значений, поля структуры инициализируются сами
	undefined := person3{}
	fmt.Println(undefined)

	//Обращение к полям структры и их изменение происходит через точечную нотацию "."
	alice.age = 34
	fmt.Println("Имя : ", alice.name)
	fmt.Println("Компания : ", alice.company)
	fmt.Println("Возраст : ", alice.age)

	//Необязательно объявлять новый тип при объявлении структур
	//можно использовать анонимные структуры
	/*
		var tom struct{
			field type
			field type
			field type
			...
		}
	*/
	var tom2 struct {
		field1 int
		field2 string
		field3 bool
	}

	//обращение к анонимной структуре tom
	tom.age = 2
	tom.company = "Microsoft"
	tom.name = "Anonymous Tom"

	//обращение к анонимной структуре tom2
	tom2.field1 = 2
	tom2.field2 = "Microsoft"
	tom2.field3 = false

	//Анонимные структуры можно создавать вот таким способом
	var tom5 = struct {
		name string
		age  int
	}{ //Инициализация структуры
		name: "Tom",
		age:  9,
	}
	fmt.Println(tom5.age)
	fmt.Println(tom5.name)

	tom6.age = 101
	tom6.name = "dsds"
	fmt.Println(tom6.age)
	fmt.Println(tom6.name)

	// инициализация структуры с анонимными полями
	var anonStruct = anon{bool: true, string: "SomeString", int8: 100, int16: 23232, int: 2333124234, rune: 232}
	fmt.Println(anonStruct.bool)
	fmt.Println(anonStruct.int)
	fmt.Println(anonStruct)

	// Если имеется два поля или более полей одного типа данных, то для остальных
	// нужно указать имена полей
	var anonymStructure struct {
		string
		bool
		secondString string // Иначе string redeclared - other declaration of string
	}
	anonymStructure.bool = false
	anonymStructure.string = "String"
	anonymStructure.secondString = "SecondString"
	fmt.Println(anonymStructure)

	// Как и в случае с обычными переменными, можно создавать указатели для структур

	type somePerson struct {
		name       string
		age        int
		secondName string
	}
	var tom7 *somePerson = &somePerson{"Tom", 38, "Smith"} // Указатель на объект Tom
	fmt.Println(tom7)

	//Чтобы обратиться к полям структуры через указатель, используется точечная нотация "."

	fmt.Println(tom7.age)
	fmt.Println(tom7.name)
	fmt.Println(tom7.secondName)

	tom7.age = 23
	tom7.name = "Thomas"
	tom7.secondName = "Winstone"
	fmt.Println(tom7.age)
	fmt.Println(tom7.name)
	fmt.Println(tom7.secondName)

	//Для обращения также можно использовать и размынование указателя
	(*tom7).age = 45
	(*tom7).name = "NoThomas"
	(*tom7).secondName = "Willam"
	fmt.Println(tom7.age)
	fmt.Println(tom7.name)
	fmt.Println(tom7.secondName)

	//Структуру можно создать с помощью функции new(type)
	var somePersonPtr *somePerson = new(somePerson)
	somePersonPtr.age = 9203
	somePersonPtr.secondName = "Garry"
	somePersonPtr.name = "Linux"

	fmt.Println(somePersonPtr)

	//Создание анонимной структуры с указателем на неё
	bob := new(struct {
		name, company string
		age           int
	})
	fmt.Println(bob)

	bob.name = "Bob"
	bob.company = "SunMicrosystems"
	bob.age = 47
	fmt.Println(bob)

	//Копирование структур
	tom8 := somePerson{"Tom8", 24, "Moscow center sparc technology"}
	tom9 := tom8
	fmt.Println(tom8)
	fmt.Println(tom9)

	tom9.secondName = "Barsicov"
	fmt.Println(tom8)
	fmt.Println(tom9)

	var peson2 somePerson2 = somePerson2{"SomeName", 34, "BarsikSoft"}
	increment_age_for_somePeople(&peson2)
	fmt.Println(peson2)

	//сравнение структур
	//Одноименные структуры равны тогда и только тогда, когда у них одинаковы поля как по количеству, так и по типу
	var somePerson_tom somePerson2 = somePerson2{"Tom", 56, "BarsikSoft"}
	var bob1 somePerson2 = somePerson2{"Bob", 56, "BarsikSoft"}
	var Bob2 somePerson2 = somePerson2{"Bob", 56, "BarsikSoft"}
	fmt.Println("somePerson_tom == bob1 : ", somePerson_tom == bob1)
	fmt.Println("Bob2 == bob1 : ", Bob2 == bob1)
}

// Структуры можно передавать в функции
func increment_age_for_somePeople(person *somePerson2) {
	person.age++
	fmt.Println(*person)
}

type somePerson2 struct {
	name       string
	age        int
	secondName string
}
