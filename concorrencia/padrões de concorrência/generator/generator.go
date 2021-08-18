package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Olá mundo"
	canal := escrever(msg)
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Texto recebido %s", texto)
			time.Sleep(time.Second)
		}
	}()

	return canal
}
