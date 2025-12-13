package main

import "fmt"

func main() {
	fmt.Println("Привет, ALT Linux!")
	var version string = goVersion()
	fmt.Printf("Версия Go: %s\n", version)
}

func goVersion() string {
	return "1.24.10"
}
