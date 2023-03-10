package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// METODO -> vinculado a uma struct
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())
	usuario1.fazerAniversario()
	fmt.Println(usuario1)
}
