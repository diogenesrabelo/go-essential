package main

import (
	"fmt"
	"testes-automatizados/endereco"
)

func main() {
	tipoEndereco := endereco.TipoEndereco("Largo São Pedro")
	fmt.Println(tipoEndereco)
}
