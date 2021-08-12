package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numeros := range numeros {
		total += numeros
	}
	return total
}

func main() {
	soma := soma(1, 2, 3, 20, 12, 1245, 234)
	fmt.Println(soma)
}
