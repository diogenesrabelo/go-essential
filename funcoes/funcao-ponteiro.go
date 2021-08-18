package main

import "fmt"

func modulo(n1 *int) *int {
	if *n1 >= 0 {
		return n1
	} else {
		*n1 = *n1 * -1
		return n1
	}
}

func subtracao() {

}

func main() {
	n1 := -10
	fmt.Println("Antes", n1)
	fmt.Println(modulo(&n1))
	fmt.Println("Depois", n1)
}
