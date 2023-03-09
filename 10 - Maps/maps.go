package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println("---------")
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])
	fmt.Println("---------")

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Ciencias da computacao",
			"campus": "UNICID",
		},
	}

	delete(usuario2, "curso")

	usuario2["signo"] = map[string]string{
		"nome": "Escorpiao",
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"])
	fmt.Println(usuario2["nome"]["ultimo"])
	fmt.Println("---------")
}
