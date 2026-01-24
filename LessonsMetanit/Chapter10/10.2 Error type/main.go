package main

import (
	"errors" // пакет errors уже имеет реализованные функции управления ошибки.
	"fmt"
)

type param_error struct{

}


func factorial(n int) (int, error){
    if n < 0{
        //return 0, param_error{}
        return  0, errors.New("Недопустимое число: должно быть положительным")
    }
    result := 1
    for i := 1;i <= n;i++{
        result *= i
    }
    return result, nil
}


func (error_object param_error) Error() string{
    return "Invalid parameter"
}
func main(){
    obj := param_error{}
    fmt.Println(obj.Error())
    fmt.Println(param_error{})


    // Корректный параметр
    fact, err := factorial(5)
    fmt.Println("Factorial:", fact)
    fmt.Println("Error:", err)

    fmt.Println()

    fact, err = factorial(-5)
    fmt.Println("Factorial:", fact)
    fmt.Println("Error:", err)
}