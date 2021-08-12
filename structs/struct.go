package main

import "fmt"

type user struct {
	name    string
	year    uint8
	andress andress
}

type andress struct {
	street string
	number uint16
}

func main() {
	var endereco andress
	endereco.number = 56
	endereco.street = "Rua São João"
	var usuario user
	usuario.name = "Victor"
	usuario.year = 27
	usuario.andress = endereco
	fmt.Println(usuario)

	usuario2 := user{"Diogenes", 27, endereco}
	fmt.Println(usuario2)

	usuario3 := user{year: 26}

	fmt.Println(usuario3)

}
