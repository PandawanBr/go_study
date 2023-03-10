package main

import "fmt"

func closure() func() {
	// texto := "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	// Referenciam variaveis fora do corpo
	fmt.Println("Closure")

	texto := "Dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
