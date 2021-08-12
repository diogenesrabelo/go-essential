package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
	sexo  string
}

type estudante struct {
	pessoa
	curso    string
	semestre uint8
}

func main() {
	fmt.Println("Herança")
}
