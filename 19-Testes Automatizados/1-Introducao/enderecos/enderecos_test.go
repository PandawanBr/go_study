package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// TESTE UNITARIO
	// Começa com 'Test' e o nome da função

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Estrada qualquer", "Estrada"},
		{"RUA DOS BOBOS ABC", "Rua"},
		{"Praça das Rosas", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado)
		}
	}
}

/*
go test -> executa os testes que termina com _test e começa o metodo com Test
go test ./... -> fala pro go entrar em todos os pacotes e executar os testes de todo mundo
go test -v -> mais descritivo / roda em paralelo
go test --cover -> coverage
go test --coverprofile nomearquivo.txt -> gera um arquivo txt
go tool cover --func=nomearquivo.txt -> le o arquivo e apresenta mais visivel no terminal
go tool cover --html=nomearquivo.txt -> le o arquivo e apresenta mais detalhado e descrito no navegador

o package pode mudar o nome, se manter o mesmo do pacote adicionando _test
necessario fazer o import das funcoes do pacote
*/
