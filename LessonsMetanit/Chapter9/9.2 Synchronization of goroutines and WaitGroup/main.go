package main

import (
	"fmt"
	"sync"
	"time"
)

func goruting_work(id int, wg *sync.WaitGroup){
    fmt.Printf("Горутина %d начало выполнения \n", id)
    time.Sleep(2 * time.Second)
    fmt.Printf("Горутина %d завершила выполнение \n", id)
    wg.Done() // Сигнализируем, что горутина завершила работу
}

func main(){
    var wg sync.WaitGroup
    wg.Add(2) // В группе две горутины

    fmt.Println(wg)

    var work func(int) = func(id int){
        fmt.Printf("Горутина %d начало выполнения \n", id)
        time.Sleep(2 * time.Second)
        fmt.Printf("Горутина %d завершила выполнение \n", id)
        wg.Done() // Сигнализируем, что горутина завершила работу
    }

    go work(1)
    go work(2)
    wg.Wait() // Ожидаем завершение обеих горутин

    fmt.Println("Горутины завершили выполнение")

    // Чтобы запустить в горутине функцию, которая не является анонимной туда необходимо передать указатель на
    // WaitGroup

    wg.Add(2)
    go goruting_work(1, &wg)
    go goruting_work(2, &wg)

    wg.Wait()
    fmt.Println("Горутины завершили выполнение")
}