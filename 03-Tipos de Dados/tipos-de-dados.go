package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64, int (usa arquitetura do pc 32bits/64bits)
	var numero1 int16 = 100
	fmt.Println(numero1)

	// unsigned int
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias -> rune = int32
	var numero3 rune = 123456
	fmt.Println(numero3)

	// alias -> byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	var str string = "LALALALALALA"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	texto = "teste"
	fmt.Println(texto)

	var booleano1 bool = true
	var booleano2 bool

	fmt.Println(booleano1, booleano2)

	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}
