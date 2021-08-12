package main

import (
	"fmt"
	"time"
)

func main() {
	idade := 18

	if idade >= 18 {
		fmt.Println("Maior de 18 anos, liberado!")
	} else {
		fmt.Println("Menor de 18 anos, barrado!")
	}

	if maiorIdade := idade; idade >= 18 {
		fmt.Println(maiorIdade, ", maior de 18 anos, liberado!")
	} else {
		fmt.Println(maiorIdade, ", menor de 18 anos, barrado!")
	}

	switch {
	case idade >= 18:
		fmt.Println("Maior de idade, pode entrar!")
	case idade >= 12:
		fmt.Println("Liberado para o matinê!")
	default:
		fmt.Println("Vai pra casa, moleque!")
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In three days.")
	default:
		fmt.Println("Too far away.")
	}
}
