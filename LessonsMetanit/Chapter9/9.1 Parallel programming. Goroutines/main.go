package main

import (
	"fmt"
	"time"
)

/*
   Для определения горутин применяется оператор go, который ставится перед вызовом функции:
   go вызов_функции

*/
func main(){
    for i := 1;i < 7;i++{
        go sum(i)
        // Горутины могут выполнять анонимные функции
        go func(i int){
            result := 0
            for j:=1; j < i + 1; j++{
                result += j
            }
            fmt.Println("Anonimus function gorutine:", result)
        }(i)
    }
    time.Sleep(2000 * time.Millisecond)
    fmt.Println("The End")
}

func sum(n int){
    fmt.Print("Gorutine : ", n)
    result := 0
    for i := 1; i <= n; i++{
        result += i
    }
    fmt.Println()
}