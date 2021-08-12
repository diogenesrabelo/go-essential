package main

import "fmt"

func somaAndSub(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, sub := somaAndSub(10, 5)

	fmt.Println(soma, sub)
}
