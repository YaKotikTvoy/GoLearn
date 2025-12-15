package main

import (
	"fmt"
)

func main() {
	var _int8 int8 = 127
	var _int8_2 int8 = -128
	fmt.Println(_int8, _int8_2)

	var _int16 int16 = 32767
	var _int16_2 int16 = -32768
	fmt.Println(_int16, _int16_2)

	var _int32 int32 = 2_147_483_647
	var _int32_2 int32 = -2_147_483_648
	fmt.Println(_int32, _int32_2)

	var _int64 int64 = 9_223_372_036_854_775_807
	var _int64_2 int64 = -9_223_372_036_854_775_808
	fmt.Println(_int64, _int64_2)

	var _uint8 uint8 = 255
	fmt.Println(_uint8)

	var _uint16 uint16 = 65_535
	fmt.Println(_uint16)

	var _uint32 uint32 = 4_294_967_295
	fmt.Println(_uint32)

	var _uint64 uint64 = 18_446_744_073_709_551_615
	fmt.Println(_uint64)

	//byte - синоним слова uint8
	var _byte byte = 0x_0
	var _byte_2 byte = 0x_ff
	fmt.Println(_byte, _byte_2)

	//rune - синоним для int32
	var _rune rune = 0x_fffffff
	fmt.Println(_rune)

	// Числа с плавающей запятой
	var _float32 float32 = 18
	var _float64 float64 = 18.002e10
	fmt.Println(_float32, _float64)

	//Комплексные числа
	var _complex1 complex64 = 1 + 2i
	var _complex2 complex128 = 120302302 + (-223223i)
	fmt.Println(_complex1, _complex2)

	//bool
	var isAlive bool = true
	var isEnabled bool = false

	fmt.Println(isAlive, isEnabled)

	//Строки
	var name string = "Hello go\r\t\nsome text \" \\"
	fmt.Println(name)

	//Многострочная строка
	name =
		`
	фвфвдлфджвлфывждфвфджвлфджв
	дфжвдфыжвджфджвдфывдфыджвфджывждфв
	`
	fmt.Println(name)

	//В Go нет переменных, которые имели бы null при не задании типизации
	//Все типы имеют значения по умолчанию
	//Например
	var int___ int         // 0
	var some_string string // "" - пустая строка
	var bool_ bool
	fmt.Println(int___, some_string, bool_)

	//Implicit typization
	var someString = "sadasd"
	fmt.Println(someString)

	//Краткое определение переменной, доступной только в функциях
	name23 := ""
	fmt.Println(name23)

	//Узнать тип переменной с помощью метода fmt.Printf() и спецификатора "%T"
	age := 41
	isEnabled2 := false
	message := "Hello"
	koef := 1.2e-10

	fmt.Printf("age type: %T\n", age)
	fmt.Printf("isEnabled2 type: %T\n", isEnabled2)
	fmt.Printf("message type: %T\n", message)
	fmt.Printf("koef type: %T\n", koef)

	var (
		name5 = "Tom"
		age5  = 27
	)
	fmt.Println(name5, age5)

	// А ещё, чтобы сделать отладку в VS Code код на Go нужно
	// $ go mod init 'Название пакета'.
}
