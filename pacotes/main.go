package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("dvmrabelo@gmail.com")
	fmt.Println(erro)
}
