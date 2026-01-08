package main

import (
	"fmt"
	"strings"
)

/*
   В Go для работы со строками имеется целый пакет, который предоставляет дополнительные функции для работы со
   строками в UTF-8 - "string"

   ToUpper(): преобразует каждый символ строки в верхний регистр

   ToLower(): используется для преобразования каждого символа строки в нижний регистр.

   HasPrefix(): проверяет, начинается ли строка с указанной строки или нет

   HasSuffix(): проверяет, заканчивается ли строка указанной строкой.

   Contains(): проверяет, содержит ли строка определенную подстроку

   ContainsAny(): проверяет, имеется какой-либо символ строки в другой строке

   Count(): сообщает, сколько раз определенная подстрока встречается в строке

   Join(): объединяет все элементы массива в одну строку.

   Replace(): заменяет в строку одну подстроки на другую

   ReplaceAll: заменяет все вхождения старой подстроки на новую

   Split(): возвращает массив подстрок из строки, разделенных заданным разделителем

   Trim / TrimLeft / TrimRight: удаляет определенные символы в начале и (или) в конце строки

   TrimFunc / TrimLeftFunc / TrimRightFunc: удаляет символы, которые соответствуют определенному условию, в начале и (или) в конце строки

   TrimSpace: удаляет все начальные и конечные пробелы.

   TrimPrefix: удаляет подстроку из начала строки.

   TrimSuffix: удаляет подстроку из конца строки.

   Index(): возвращает индекс первого вхождения подстроки в строке.

   LastIndex(): ищет последнее вхождение подстроки в строку

   IndexAny(): возвращает первый индекс любого из найденных символов подстроки в строке.

   LastIndexAny(): возвращает последнее вхождение символа в строке.
*/


func main(){
    // Преобразование строки в верхний и нижний регистр
    str := "Hello world"

    fmt.Println("___________________ToUpper__________________________")
    fmt.Println(strings.ToUpper(str))

    fmt.Println("___________________ToLower__________________________")
    fmt.Println(strings.ToLower(str))

    // Проверка начинается ли строка на какую-либо последовательность - HasPrefix(str, prefix)
    // или наоборот заканчивается на какую либо последовательность - HasSuffix(str, suffix)

    fmt.Println("___________________HasPrefix________________________")
    fmt.Println("Start str with \"He\":", strings.HasPrefix(str, "He")) // true
    fmt.Println("Start str with \"His\":", strings.HasPrefix(str, "His")) // false


    fmt.Println("___________________HasSuffix________________________")
    fmt.Println("Finished str with \"ld\":", strings.HasSuffix(str, "ld")) // true
    fmt.Println("Finished str with \"old\":", strings.HasSuffix(str, "old")) // false

    // Чтобы точно проверить наличие последовательности элементов в срезе
    // используется string.Contains()

    // Если нужно проверить просто наличие хотя бы одного элемента из поданной строки
    // не взирая на строгую последовательность, то используется string.ContainsAny()
    fmt.Println("___________________Contains_________________________")
    fmt.Println("Contains \"world\" :", strings.Contains(str, "world"))
    fmt.Println("Contains \"work\" :", strings.Contains(str, "work"))

    // Чтобы просто узнать есть ли в строке хотя бы один символ из последовательности
    fmt.Println("___________________ContainsAny______________________")
    fmt.Println("ContainsAny \"sam\" :", strings.Contains(str, "sam")) // false
    fmt.Println("Contains \"work\" :", strings.Contains(str, "work")) // true // есть символ 'w' 'o' 'r' в строке "Hello world"

    // Count() - возвращает количество вхождений строки в какой-либо другой строке
    fmt.Println("___________________Count____________________________")
    fmt.Println("Count of \"world\" :", strings.Count(str, "world"))
    fmt.Println("Count of \"o\" :", strings.Count(str, "o"))
    fmt.Println("Count of \"foo\" :", strings.Count(str, "foo"))

    // Join([]string, divider string) - возвращает строку, которая будет объединена из среза строк с указаным
    // разделителем
    fmt.Println("___________________Join_____________________________")
    var Join_massive []string = []string{"Hello", "world", "!", "!"}
    fmt.Println(strings.Join(Join_massive, " "))

    // Замена подстроки с помощью метода Replace, заменяет вхождения
    fmt.Println("___________________Replace__________________________")
    str = "Hello world, good bye world"
    fmt.Println(str)
    fmt.Println(strings.Replace(str, "world", "work", 1))
    fmt.Println(str)

    // strings.ReplaceAll() - заменяет все вхождения строки на строку
    fmt.Println("___________________ReplaceAll_______________________")
    fmt.Println(strings.ReplaceAll(str, "world", "work"))

    str = "+7-987-654-32-10"
    fmt.Println("___________________Split____________________________")
    fmt.Println(str)
    fmt.Println(strings.Split(str, "-"))


    fmt.Println("___________________Field___________________________")
    str = "Hello World\nGood bye World\n"
    // разбиваем строку по пробелам или переводам строки \n
    fmt.Println(str)
    fmt.Println(strings.Fields(str))


    fmt.Println("___________________Trims___________________________")
    fmt.Println(strings.Trim("01.2300", "0")) // Удаление справа и слева //1.23
    fmt.Println(strings.TrimLeft("01.2300", "0")) // Удаление  слева //1.2300
    fmt.Println(strings.TrimRight("01.2300", "0")) // Удаление  слева //01.23


    str = "+7-987-654-32-10!"
    // Удаление справа и слева по условию,
    // оставляем числа - в кодировке ASCII числа
    // числа кодируются кодом от a >= '0' && a <= '9', всё остальное - знакописные символы //1.23
    fmt.Println(strings.TrimFunc(str, func(c rune) bool {return c < '0' || c > '9'})) //79876543210
    fmt.Println(strings.TrimRightFunc(str, func(c rune) bool {return c < '0' || c > '9'})) //+7-987-654-32-10
    fmt.Println(strings.TrimLeftFunc(str, func(c rune) bool {return c < '0' || c > '9'})) //7-987-654-32-10!

    // Обрезка префикса
    fmt.Println(strings.TrimPrefix(str, "+7")) //-987-654-32-10!
    // Обрезка суффикса
    fmt.Println(strings.TrimSuffix(str, "!")) //+7-987-654-32-10
    // Обрезка начальных и конечных пробелов
    fmt.Println(strings.TrimSpace(" 7 987 654 32 10 ")) // "7 987 654 32 10"




    str = "Hello world"
    // Ищет индекс первого вхождения в строку
    fmt.Println("___________________Index___________________________")
    fmt.Println("Index \"He\" :", strings.Index(str, "He")) // 0
    fmt.Println("Index \"lo\" :", strings.Index(str, "lo")) // 3
    fmt.Println("Index \"f\" :", strings.Index(str, "f")) // -1

    str = "Hello world. Good bye world."
    // Ищет индекс последнего вхождения в строку
    fmt.Println("___________________LastIndex_______________________")
    fmt.Println("LastIndex \"ld\" :", strings.LastIndex(str, "ld")) // 0
    fmt.Println("LastIndex \"lo\" :", strings.LastIndex(str, "lo")) // 3
    fmt.Println("LastIndex \"f\" :", strings.LastIndex(str, "f")) // -1

    str = "Hello world. Good bye world."
    // Ищет индекс первого вхождения в строку любого символа из строки
    fmt.Println("___________________IndexAny________________________")
    fmt.Println("IndexAny \"ld\" :", strings.IndexAny(str, "ld")) // 2
    fmt.Println("IndexAny \"Hsd\" :", strings.IndexAny(str, "Hsd")) // 0
    fmt.Println("IndexAny \"f\" :", strings.IndexAny(str, "f")) // -1


    str = "Hello world. Good bye world."
    // Ищет индекс первого вхождения в строку любого символа из строки
    fmt.Println("_______________LastIndexAny________________________")
    fmt.Println("LastIndexAny \"ld\" :", strings.LastIndexAny(str, "ld")) // 26
    fmt.Println("LastIndexAny \"Hsd\" :", strings.LastIndexAny(str, "Hsd")) // 26
    fmt.Println("LastIndexAny \"fo\" :", strings.LastIndexAny(str, "fo")) // 23
}
