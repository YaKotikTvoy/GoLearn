package main

/*
Гонка данных (data race) представляет ситуацию, когда два или несколько потоков одновременно обращаются к одному и тому
же участку памяти и выполняют по крайней мере одну операцию записи. Например:
*/
import (
	"fmt"
	"sync"
	"time"
)
var value int = 22

var wg sync.WaitGroup

func update_value (id int){
    fmt.Println("Goroutine", id, "starts")
    if value == 22{
        time.Sleep(1 * time.Second)
        fmt.Println("Goroutine", id, "changes the value")
        value += 1
    }
    fmt.Println(value)
    wg.Done() //Сигнализируем, что горутина завершила выполнение
}

func main(){
    wg.Add(2) // Ждем две горутины

    go update_value(1)
    go update_value(2)

    wg.Wait()
}


// Для обнаружения гонок необходимо компилировать или запускать программу с флагом -race
// $ go run -race file.go
// $ go build - race file.go