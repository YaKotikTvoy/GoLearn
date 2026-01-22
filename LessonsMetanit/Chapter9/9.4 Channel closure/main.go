package main

import "fmt"


func send_value(c chan int){
    c <- 10 // Отправляем горутине main число 10


    // После закрытия канала в него нельзя посылать новые данные
    // и при попытке получить из него данные, которые уже вычитаны, там отдастся значение поумолчанию
    close(c)
    //c <- 22 // Ошибка канал уже закрыт
}

func send_many_value(c chan int){
    // Небуферизированный канал
    c <- 1
    c <- 2
    c <- 3
    c <- 4
    c <- 5
    c <- 6
    close(c)
}

func get_value(c chan int){

    // Чтобы не столкнуться со случаем, когда канал уже закрыт, можно проверять состояние канала,
    // это делается с помощью возвращения второго значения

    val, ok := <-c
    if ok {
        fmt.Println("Канал открыт, значение: ", val)
    } else {
        fmt.Println("Канал закрыт")
    }
}
func main(){
    var intCh chan int = make (chan int)
    go send_value(intCh)

    //Получаем данные
    get_value(intCh) // 10
    get_value(intCh) // 0

    var intCh2 chan int = make(chan int)
    go send_many_value(intCh2)

    for true{
        value, ok := <- intCh2
        if !ok{
            break
        }
        fmt.Println("Получено число:", value)
    }
    fmt.Println("Канал закрыт")


    //Тоже самое относиться и к буферизированным каналам
    var intCh3 chan int = make(chan int, 3)
    intCh3 <- 8
    intCh3 <- 9
    intCh3 <- 10
    //intCh3 <- 11
    close(intCh3)
    for i:= 0; i< cap(intCh3);i++{
        if val, ok := <-intCh3; ok{
            fmt.Println(val)
        } else{
            fmt.Println("Канал закрыт")
        }
    }
}
