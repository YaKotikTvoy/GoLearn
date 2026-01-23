package main

import "fmt"

// Оператор select определяет операцию канала. Данный оператор имеет блоки case как в switch, только
// case определяет операции, которые произойдут с каналом.

/*
select{

    case операция_1:
        // обработка операции_1
    case операция_2:
        // обработка операции_2
..........................................
    case операция_N:
        // обработка операции_N
}
*/

func square(c chan int){
    value := <- c
    c <- value * value
}

func cube(c chan int){
    value := <- c
    c <- value * value * value
}

func double(c chan int){
    value := <- c
    c <- value + value
}

func main(){

    sqr_ch := make ( chan int)
    cube_ch := make ( chan int)

    // Запускаем две горутины для взаимодействия через канал
    go square(sqr_ch)
    go cube(cube_ch)

    // Отправляем данные в канал
    sqr_ch <- 2
    cube_ch <- 2

    //Обработка данных
    select{
        case sqr_val := <- sqr_ch:
            fmt.Println("Square:", sqr_val)
        case cube_val := <- cube_ch:
            fmt.Println("Cube:", cube_val)
        default: // Если в операторе case нет соответствующего блока case, которые могут выполнить операцию, то это
        // приведёт к взаимоблокировке. Чтобы избежать блокировки нужен блок, выполниющийся в любом случае

    }



    //Обработка данных
    select{
        case sqr_val := <- sqr_ch:
            fmt.Println("Square:", sqr_val)
        case cube_val := <- cube_ch:
            fmt.Println("Cube:", cube_val)
        default:
            fmt.Println("Undefined operation.")

    }

    // Select пройдется только по результатом одной отработанной горутины
    // Чтобы обработать результаты всех горутин, нужно select поместить в цикл

    square_ch1 := make(chan int)
    cube_ch1 := make(chan int)
    double_ch1 := make(chan int)

    go square(square_ch1)
    go cube(cube_ch1)
    go double(double_ch1)
    square_ch1 <- 3
    cube_ch1 <- 3
    double_ch1 <- 3


    for range 3{
        select{
            case square_r := <-square_ch1:
                fmt.Println("Квадрат:", square_r)
            case cube_r := <-cube_ch1:
                fmt.Println("Куб:", cube_r)
            case double_r := <- double_ch1:
                fmt.Println("Сумма:", double_r)
        }
    }

}