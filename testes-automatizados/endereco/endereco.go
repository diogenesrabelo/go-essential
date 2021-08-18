package endereco

import "strings"

//TipoEndereco verifica o tipo de enredeço
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "largo", "estrada"}

	primeiraPalavra := strings.ToLower(strings.Split(endereco, " ")[0])

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			return strings.Title(tipo)
		}
	}
	return "Desconhecido"
}
