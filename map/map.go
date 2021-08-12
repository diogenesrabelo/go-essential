package main

import "fmt"

func main() {
	fmt.Println("Map")

	usuario := map[string]string{
		"nome":      "Diogenes",
		"sobrenome": "Rabelo",
		"idade":     "27",
	}

	fmt.Println(usuario)
	fmt.Println("Nome: ", usuario["nome"])
	fmt.Println("Sobrenome: ", usuario["sobrenome"])
	fmt.Println("Idade: ", usuario["idade"])
}
