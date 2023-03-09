package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario

	fmt.Println(u)

	u.nome = "Rodrigo"
	u.idade = 27

	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}
	usuario2 := usuario{"Gabriel", 24, enderecoExemplo}

	fmt.Println(usuario2)

	usuario3 := usuario{idade: 20}
	fmt.Println(usuario3)
}
