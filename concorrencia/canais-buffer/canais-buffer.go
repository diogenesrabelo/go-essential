package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olé mundo"
	canal <- "Canal com buffer"
	// canal <- "Canal com buffer" descomentar causa deadlock!
	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
