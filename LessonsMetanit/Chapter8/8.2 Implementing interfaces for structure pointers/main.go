package main

import "fmt"

type Reader interface{
    read()
}

type Writer interface{
    write(string)
}

type File struct{
    text string
}

// Реализация интерфейса Reader для File
//func (f File) read(){
//    fmt.Println(f.text)
//}


// Тем не менее в Go можно реализовывать
// реализацию интерфейса непосредственно для указателя на структру
func (f *File) read(){
    fmt.Println(f.text)
}



func (f *File) Write(message string){
    f.text = message
    fmt.Println("В файл записано", message)
}
/*
func (f File) Write(message string){
    f.text = message
    fmt.Println("В файл записано", message)
}
*/

func write_data(data Writer, message string){
    data.write(message)
}

func read_data(data Reader){
    data.read()
}

func main(){
    fmt.Println()

    var file File = File{"Какой-то файл"}

    read_data(&file)

    var p_file *File = new(File)
    p_file.text = "Ещё какой-то файл"

    // в методы, параметры которых реализуют интерфейсы можно помещать указатели типов, которые реализуют эти интерфейсы
    read_data(p_file) // И так можно

    file = File{"Undefined"}
    file.Write("Я вписал какой-то текст в файл")
    fmt.Println(file.text) // Undefined, если интерфейс реализован к типу приемнику без указателя
}