package main

import "fmt"

// Бывают такие ситуации, что в в канал одна горутина отправляет целый поток данных.
// В этом случае горутина-получатель, принимает поток в бесконечном цикле,
// и если получен маркер закрытия канала, то происходит выход из цикла.

func main(){
    intCh := make(chan int)

    go factorial(7, intCh)

    for{
        var num, opened = <- intCh
        if !opened {
            break
        }
        fmt.Println(num)
    }

    // Также такой канал с потоком данных, можно пробежать аналогом foreach в Go
    intCh2 := make(chan int)
    go factorial(10, intCh2)


    fmt.Println()
    // То есть как обычный срез
    for i := range intCh2{
        fmt.Println(i)
    }
}

func factorial(n int, ch chan  int){
    result := 1
    for i := 1; i <= n; i++{
        result *= i
        ch <- result // посылаем по числу
    }
    defer close(ch) // Оператор, который закрывает канал по завершению вычислений функции
}