package main

import "fmt"

func diaDaSemana(numero int) string {

	var diaDaSemana string

	switch numero {
	case 1:
		diaDaSemana = "Domingo"
		// fallthrough cai na proxia condicao sem validar
		fallthrough
	case 2:
		diaDaSemana = "Segunda-Feira"
	case 3:
		diaDaSemana = "Terça-Feira"
	case 4:
		diaDaSemana = "Quarta-Feira"
	case 5:
		diaDaSemana = "Quinta-Feira"
	case 6:
		diaDaSemana = "Sexta-Feira"
	case 7:
		diaDaSemana = "Sábado-Feira"
	default:
		diaDaSemana = "Número inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(8))
}
