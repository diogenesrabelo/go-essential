package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type pessoa struct {
	Nome  string `json:"nome"`
	Idade uint8  `json:"idade"`
	Sexo  string `json:"sexo"`
}

func main() {
	pessoa1 := pessoa{"Diogenes", 27, "masculino"}

	pessoaEmJSON, erro := json.Marshal(pessoa1)
	if erro != nil {
		log.Fatal("Erro ao converter pessoa em JSON")
	}

	fmt.Println(bytes.NewBuffer(pessoaEmJSON))

	pessoa2 := map[string]string{
		"nome":  "Victor",
		"idade": "24",
		"sexo":  "masculino",
	}

	pessoa2EmBytes, erro := json.Marshal(pessoa2)
	if erro != nil {
		log.Fatal("Erro ao converter pessoa em JSON")
	}
	pessoa2EmJSON := bytes.NewBuffer(pessoa2EmBytes)
	fmt.Println(pessoa2EmJSON)

	var p1 pessoa
	if erro := json.Unmarshal(pessoaEmJSON, &p1); erro != nil {
		log.Fatal("Erro ao converter JSON em Objeto")
	}

	fmt.Println(p1)

	p2 := make(map[string]string)
	if erro := json.Unmarshal(pessoa2EmBytes, &p2); erro != nil {
		log.Fatal("Erro ao converter JSON em Objeto")
	}
	fmt.Println(p2)
}
