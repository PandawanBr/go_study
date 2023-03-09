package main

import "fmt"

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A MEDIA É EXATAMENTE 6!")
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso")
	}
}

func main() {
	// PANIC -> PARA TODO EXECUCAO DO PROGRAMA -> EXECUTA TODOS OS DEFER ANTES
	// RECOVER -> RECUPERA EXECUCAO DEPOIS DO PANIC
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execucao")
}
