package main

import (
	"fmt"
)

/*
Интерфейс представляет собой абстракцию, содержащий набор методов, которые нужно реализовать.

type имя_интерфейса interface{
    определения_функций
}
*/

type Vehicle interface{
    move()
}


type Car struct {}
type Aircraft struct {}
func (c Car) move(){
    fmt.Println("Car drives")
}

func (c Aircraft) move(){
    fmt.Println("Aircraft drives")
}

func moveVehicle(v Vehicle){
    v.move()
}

func moveAircraft(a Aircraft){
    a.move()
}

func moveCar(c Car){
    c.move()
}


type Movable interface{
    MoveX(distance int)
    MoveY(distance int)
}

type Rectangle struct{

}


func (r Rectangle) MoveX(distance int){
    fmt.Println("Перемещаем прямоугльник на",distance, "см по X" )
}

func (r Rectangle) MoveY(distance int){
    fmt.Println("Перемещаем прямоугльник на",distance, "см по Y" )
}

func Move_object(r Rectangle){
    r.MoveX(10)
    r.MoveY(-10)
}



// пустой интерфейс, используют когда не уверены к какому типу относится интерфейс
type Empty interface{ }

func print_empty(e Empty){
    fmt.Println(e)
}

type person struct{
    name string
}

type account struct{
    name string
}

type car1 struct{
    model string
}
type aircraft1 struct{
    model string
}

func (c car1) move(){
    fmt.Println(c.model, "едет")
}

func (c aircraft1) move(){
    fmt.Println(c.model, "летит")
}

func main(){
    var v Vehicle
    fmt.Println(v) // nil
    v = Car{}
    v.move()
    var v2 Vehicle
    v2 = Aircraft{}
    v2.move()

    var car Car = Car{}
    moveCar(car)
    var aircraft Aircraft = Aircraft{}
    moveAircraft(aircraft)

    moveVehicle(car)
    moveVehicle(aircraft)

    var r Rectangle = Rectangle{}
    Move_object(r)


    var person person = person{name: "Tom"}
    var account account = account{name: "Tom account"}

    print_empty(person)
    print_empty(account)

    var vehicles [3]Vehicle = [3]Vehicle{car1{"Volga"}, aircraft1{"Boeng"}, aircraft1{"Su-57"}}

    for _, v := range vehicles{
        fmt.Println(v)
        v.move()
    }
}
