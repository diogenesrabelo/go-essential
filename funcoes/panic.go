package main

import "fmt"

func recuperardoPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recurepada com sucesso!")
	}
}

func calculaMedia(n1, n2 float64) bool {
	defer recuperardoPanic()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A media é 6!")
}

func main() {
	fmt.Println(calculaMedia(6, 6))
	fmt.Println("Pós exec!")
}
