package main

import "fmt"

// 1 POR FUNCAO E OBRIGATORIAMENTE SER O ULTIMO
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(soma())
	escrever("Hello", 1, 2, 3, 4, 5, 6)
}
