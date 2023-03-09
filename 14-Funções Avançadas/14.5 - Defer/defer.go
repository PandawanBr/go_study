package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2() {
	fmt.Println("Função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada, o resultado sera calculado")
	fmt.Println("Entrando na funcao para verificar se aluno esta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {

	// DEFER = Adiar -> Adia a execucao o maximo possivel
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
