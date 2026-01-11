package main

import "fmt"

/*
  В Go есть механизм проверки типов type assertion
  t, ok := i.(T)
  интерфейс i проверяется на наличие конкретного значения типа T. Если значение есть,
  то ok - будет true, а t будет иметь значение типа T.

  иначе ok - будет false, а t - будет 0.
*/

type Movable interface{
    move()
}
type Rectangle struct{
    x,
    y,
    height,
    width int
}
type Circle struct{
    x,
    y,
    radius int
}

type Point struct{
    x,
    y int
}

func (c Circle) move(){
    fmt.Println("Перемещаем круг")
}
func (r Rectangle) move(){
    fmt.Println("Перемещаем прямоугольник")
}

func check(i interface{}){
    switch value := i.(type){
        case Circle:
            fmt.Println("Type : Circle, value:", value)
        case Rectangle:
            fmt.Println("Type : Rectangle, value:", value)
        default:
            fmt.Println("Type: undefined")
    }
}

func main(){
    var shape_movable Movable = Rectangle{x: 1, y: 23, width: 10, height: 14}
    shape_movable.move()

    // Проверяем результат
    var i, ok = shape_movable.(Rectangle) // Проверяем является ли данный интерфейс прямоугольником
    fmt.Println(i, ok)

    i1, ok1 := shape_movable.(Circle)
    fmt.Println(i1, ok1)

    /*
    Можно делать проверку типа, реализующего интерфейс с помощью конструкции Type switch
    switch value := i.(type){

        case T1:        // Действия, если value представляет тип T1

        case T2:        // Действия, если value представляет тип T2

        .......................................................
        case TN:        // Действия, если value представляет тип TN

        default:    // если ни один из типов в case не соответствуют v
    }
    */
    check(Rectangle{1, -12, 1223, 1})
    check(Circle{0, -12, 23})
    check(Point{5, -70})
}