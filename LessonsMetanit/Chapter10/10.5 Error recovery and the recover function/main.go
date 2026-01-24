package main

import "fmt"

/*
Для обработки ошибок, вызванных функцией panic() используется функция recover(), которая должна вызваться оператором
defer после вызова panic. Это отчасти похоже на try-catch

Вызовы происходят следующим образом:
1) panic()
2) defer
3) recover в defer
*/
func main(){
    fmt.Println(divide(4,0)) // fmt.Println не получит возвращаемого значения от divide. И всё закончится выполнением
    // строки
    fmt.Println("Program has been finished")
}

func try_catch(){
    if r:= recover(); r != nil{ // вызов recover() обработывает вызов panic, а далее выполнение возвращается
        fmt.Println("Error:", r) // на место той функции, которая вызвала panic(), то есть в функцию main
    }
}

func divide(x, y float64) float64{
    defer try_catch()
    if y == 0{
        panic("Division by zero!") // panic() останавливает выполнение программы и начинает раскручивать стек вызовов,
        // выполняя все отложенные defer-функции, которые обработают вызов panic
    }
    return x / y
}
/*

Если суммировать, то мы получим следующую последовательность действий:

Вызывается main(), затем divide(4, 0).

Внутри divide, defer try_catch() откладывает выполнение try_catch до выхода из divide.

Условие y == 0 истинно, поэтому вызывается panic("Division by zero!").

Начинается раскрутка стека. Прежде чем divide полностью завершится, срабатывает отложенная функция try_catch().

Внутри try_catch(), recover() перехватывает значение паники ("Division by zero!").

fmt.Println("Error:", r) выводит "Error: Division by zero!".

Поскольку вызов panic() был перехвачен, программа продолжает выполнение из main().

Выводится "Program has been finished".

*/