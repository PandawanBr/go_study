package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posicao 1"
	fmt.Println(array1)

	array2 := [5]string{"Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4", "Posicao 5"}
	fmt.Println(array2)

	// [...] PREENCHE COM A QUANTIDADE DE ITENS INFORMADOS NA CRIACAO DO ARRAY
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// SLICE NAO É ARRAY, MAS SE PARECE COM UM, POIS ELE REFERENCIA UM ARRAY
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	// FUNCIONA COMO UM PONTEIRO -> PRIMEIRO INDICE É INCLUSIVO, SEGUNDO INDICE É EXCLUSIVO
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posicao alterada"
	fmt.Println(slice2)

	// ARRAYS INTERNOS
	fmt.Println("-----------------")

	slice3 := make([]float32, 10, 11)
	slice3 = append(slice3, 5)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // len() para tamanho
	fmt.Println(cap(slice3)) // cap() para capacidade

	// SE NAO PASSAR A CAPACIDADE, ELE USA O TAMANHO COMO CAPACIDADE
	slice4 := make([]float32, 5)
	slice4 = append(slice4, 10)

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
