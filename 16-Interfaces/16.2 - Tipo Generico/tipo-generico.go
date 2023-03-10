package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	// Jeito errado de usar
	mapa := map[interface{}]interface{}{
		1:           "string",
		float32(10): true,
		"Teste":     1,
	}
	fmt.Println(mapa)
}
