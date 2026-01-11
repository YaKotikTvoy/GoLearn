package main

import "fmt"

// Структура может реализовывать сразу несколько интерфейсов и иметь собственные методы
//

type Movable interface{
    move()
}
type Drawable interface{
    draw()
}

type Rectangle struct{}

func (r Rectangle) draw(){
    fmt.Println("Перемещаем прямоугольник")
}

func (r Rectangle) move(){
    fmt.Println("Рисуем прямоугольник")
}

func move_object(obj Movable){
    obj.move()
}

func draw_object(obj Drawable){
    obj.draw()
}

// Go не поддерживает наследование, но позволяет интерфесам содержать в себе другие интерфейсы
//

type Writer interface {
    write(string)
}
type Reader interface {
    read()
}

type ReaderWriter interface{
    Reader
    Writer //можно, также неявно встроить интерфейс Writer
    //write(string)
}

type File struct{
    text string
}

func (f *File) read(){
    fmt.Println(f.text)
}

func (f *File) write(message string){
    f.text = message
    fmt.Println("Запись в файл строки :", f.text)
}

func WriteToStream(writer Writer, message string){
    writer.write(message)
}

func main(){
    var rect Rectangle = Rectangle{}

    move_object(rect)
    draw_object(rect)

    var draw1 Drawable = Rectangle{}
    var move1 Movable = Rectangle{}

    move_object(move1)
    draw_object(draw1)

    fmt.Println()


    var myFile ReaderWriter = new(File)
    myFile.write("Что-то, какая-то запись")
    myFile.read()

    var file *File = new(File)
    WriteToStream(file, "Какая-то строка, передаваемая объекту writer")


}