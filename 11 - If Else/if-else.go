package main

import "fmt"

func main() {
	fmt.Println("Estrutura de Controle")

	numero := 10

	// SEMPRE USA CHAVES
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// IF INIT -> INICIA A VARIAVEL DENTRO DA CONDICAO
	// VARIAVEL USADA APENAS NO ESCOPO DA CONDICAO
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Outro numero mair do que zero")
	} else if outroNumero < -10 {
		fmt.Println("Numero menos do que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

}
