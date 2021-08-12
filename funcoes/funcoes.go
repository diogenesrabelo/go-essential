package main

import (
	"fmt"
)

func main() {

	var f = func(n1 int, n2 int) int {
		if n1 > n2 {
			return n1
		}
		return n2
	}
	fmt.Println(f(2, 3))
}

func imprime(name string) {
	fmt.Println("Meu nome é ", name)
}

func soma(n1 int, n2 int) int {
	return n1 + n2
}

func sub(n1 int, n2 int) int {
	return n1 - n2
}

func mult(n1 int, n2 int) int {
	return n1 - n2
}

func div(n1 int, n2 int) int {
	return n1 - n2
}
