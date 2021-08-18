package main

import "fmt"

func main() {
	fmt.Println("Função main!")
}

func init() {
	fmt.Println("Antes da main e posso ser uma por arquivo")
}
