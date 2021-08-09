package main

import (
	"errors"
	"fmt"
)

func main() {
	var intOito int8 = 127
	var intDezesseis int16 = 32767
	var intTrintaDois int32 = 2147483647
	var intSessentaQuarto int64 = 9223372036854775807
	//Arquitetura do computador para int
	var intComputer int = 9223372036854775807
	fmt.Println(intOito)
	fmt.Println(intDezesseis)
	fmt.Println(intTrintaDois)
	fmt.Println(intSessentaQuarto)
	fmt.Println(intComputer)

	//Só positivos
	var uintOito uint8 = 255
	var uintDezesseis uint16 = 65535
	var uintTrintaDois uint32 = 4294967295
	var uintSessentaQuarto uint64 = 18446744073709551615
	var uintComputer uint = 18446744073709551615
	fmt.Println(uintOito)
	fmt.Println(uintDezesseis)
	fmt.Println(uintTrintaDois)
	fmt.Println(uintSessentaQuarto)
	fmt.Println(uintComputer)

	//Float

	var floatTrintaDois float32 = 4294967295.4
	var floatSessentaQuatro float64 = 18446744073709551615.04

	fmt.Println(floatTrintaDois)
	fmt.Println(floatSessentaQuatro)

	//String

	var name string = "Diogenes"
	fmt.Println(name)

	//Char -> Numero da tabela ascii = int
	char := 'B'
	fmt.Println(char)

	var booleano bool = true
	fmt.Println(booleano)

	var erro error = errors.New("Exemplo de erro")
	fmt.Println(erro)
}
