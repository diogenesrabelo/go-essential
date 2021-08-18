package endereco_test

import (
	"testes-automatizados/endereco"
	"testing"
)

type cenarioTest struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoEndereco(t *testing.T) {

	cenariosTest := []cenarioTest{
		{"Rua dos bobos", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Praça dos imigrantes", "Desconhecido"},
		{"Largo São Pedro", "Largo"},
		{"Estrada de itapecerica", "Estrada"},
	}

	for _, cenario := range cenariosTest {
		enderecoRecebido := endereco.TipoEndereco(cenario.enderecoInserido)
		if enderecoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado. Recebido: %s e Esperado: %s",
				enderecoRecebido,
				cenario.enderecoEsperado)
		}
	}

}
