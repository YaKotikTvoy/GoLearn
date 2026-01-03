package main

/*
Теперь используем функции пакета messages и для этого в корневой папке проекта определим главный файл приложения -
main.go со следующим кодом:
*/

/*
Для подключения пакета используется следующее выражение

import [путь к модулю]/[путь к пакету]
*/
//import "Splitting_packages/messages"

// Пакеты могут содержать другие пакеты, это зависит от вложенности папок

import (
	"Splitting_packages/messages/messages_en"
	"Splitting_packages/messages/messages_ru"
)
func main(){
    messages_en.Bye_en()
    messages_en.Hello_en()

    messages_ru.Bye_ru()
    messages_ru.Hello_ru()
}
