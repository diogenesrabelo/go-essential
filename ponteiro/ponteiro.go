package main

import "fmt"

func main() {

	var idade uint8 = 27

	var idadeCopia uint8 = idade

	fmt.Println(idade, idadeCopia)

	idade++
	fmt.Println(idadeCopia)

	var pIdade *uint8 = &idade

	fmt.Println("Endereço da variavel (valor do ponteiro) ", pIdade)
	fmt.Println("Valor apontado/referenciado pelo ponteiro ", *pIdade)
}
