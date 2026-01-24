package main

import (
	"fmt"
	"os"
)

// Чтобы сгенерировать ошибку применяется функция panic()
func main(){
    file, err := os.Open("./context.txt")
    fmt.Println("file :", file)
    fmt.Println("error:", err)

    fmt.Println()

    file, err = os.Open("~/Develop/LessonsMetanit/Chapter10/10.1 Isomorphic error handling/main.go")
    fmt.Println("file :", file)
    fmt.Println("error:", err)


    fact, err1 := factorial(5)
    fmt.Println("Factorial:", fact)
    fmt.Println("Error:", err1)

    fmt.Println()

    //Некорректный параметр
    fact, err1 = factorial(-5)
    fmt.Println("Factorial:", fact)
    fmt.Println("Error:", err1)
}

func factorial(n int)(int, interface{}){
    if n < 0{
        return 0, "Недопустимое число, должно быть больше нуля."
    }
    result := 1
    for i := 1; i <= n; i++{
        result *= i
    }
    return result, nil
}