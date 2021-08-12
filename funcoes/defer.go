package main

import "fmt"

func funcao1() {
	fmt.Println("Exec func 1")
}

func funcao2() {
	fmt.Println("Exec func 2")
}

func main() {
	//defer -> adia até o ultimo momento possível dentro do contexto/função
	defer funcao1()
	funcao2()
}
