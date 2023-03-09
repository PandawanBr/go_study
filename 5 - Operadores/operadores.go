package main

import "fmt"

func main() {
	// ARITIMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// NAO REALIZA OPERACOES ARITIMETICAS ENTRE TIPOS DIFERENTES -> int8 + int16
	var numero1 int16 = 10
	var numero2 int16 = 25
	var soma2 int16 = numero1 + int16(numero2)

	fmt.Println(soma2)

	// OPERADORES DE ATRIBUICAO
	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	// 	OPERADORES RELACIONAIS -> sempre retorna bool
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println("-----------------")
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// OPERADORES UNARIOS
	numero := 10
	numero++
	fmt.Println(numero)

	numero += 15
	fmt.Println(numero)

	numero--
	fmt.Println(numero)

	numero -= 20
	fmt.Println(numero)

	//  NAO EXISTE OPERADOR TERNARIO
	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
