package main

import "fmt"

func main() {
	fmt.Println("Loop")

	iCont := 0
	for iCont < 10 {
		fmt.Println(iCont)
		iCont++
	}

	for jCont := 0; jCont < 10; jCont++ {
		fmt.Println(jCont)
	}

	nomes := [3]string{"Diogenes", "Victor", "Rabelo"}

	// index ou _ no primeiro parametro
	for _, nome := range nomes {
		fmt.Println(nome)
	}
}
