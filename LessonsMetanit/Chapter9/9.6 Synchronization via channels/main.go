package main

import "fmt"

// Горутины общаются между собой посредством каналов, например, одна горутина вычисляет квадрат числа, вторая
// вторая принимает результат вычисления
func main(){
    var intCh chan int = make(chan int)

    go square(5, intCh)

    fmt.Println(<-intCh) // Main блокируется до тех пор пока не в канал не будет помещены данные, над которыми
    // потом необходимо будет выполнить вычисления.

    // Канал не обязательно должен нести данные, которые представляют некоторый результат, от которого зависит
    // дальнейшее выполнение горутины. Иногда это может быть холостой объект, например, пустая структура, которая
    // необходима только для синхронизации горутин:

    results := make (map[int] int)
    structCh := make (chan struct{})

    go square2(5, structCh, results)

    <-structCh // Ожидаем закрытие канала structCh

    for i, v := range results{
        fmt.Println(i,"-",v)
    }


    // Что если нужно отследить завершение нескольких горутин?
    done := make(chan bool)

    var count int = 4
    for i:=1; i <= count ; i++ {
        go sum(i+3, done)
    }

    //Ожидаем завершения всех горутин
    for true{
        v := <-done
        fmt.Println("len(done) :", len(done), "значение :", v)
        if !(len(done) > 0){
            break
        }
    }
    fmt.Println("The end")
}

func square(n int, ch chan int){
    fmt.Println("Вычисляем квадрат числа:", n)
    var result int = n * n
    ch <- result
}

func square2(n int, ch chan struct{}, results map[int]int){
    for i:=1; i <= n; i++{
        results[i] = i * i
    }
    close(ch)
}
func sum(n int, ch chan bool){
    result := 0

    for k:= 1; k <= n;k++{
        result += k
    }
    fmt.Println("Сумма первых", n, "чисел равна", result)
    ch <- true
}