package main

import "fmt"

func main() {
	texto := func(message string) string {
		return fmt.Sprintf("Recebido %s", message)
	}("Olá mundo!")
	fmt.Println(texto)
}
