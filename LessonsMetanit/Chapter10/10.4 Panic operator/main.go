package main

import "fmt"

func main(){
    defer main_defer()
    fmt.Println(divide(15.0,5.0))
    fmt.Println(divide(15.0,0x0))
    fmt.Println("Program has been finished")
}

/*
    Оператор panic() останавливает выполнение функции, а далее всей программы, если имеется несколько операторов
    defer, то они выполняться перед всей остановкой
*/

func divide(x,y float64) float64{
    defer divide_defer()
    if y == 0{
        panic("Division by zero!")
    }
    return x / y
}

func main_defer(){
    fmt.Println("main_defer executed")
}

func divide_defer(){
    fmt.Println("divide_defer executed")
}