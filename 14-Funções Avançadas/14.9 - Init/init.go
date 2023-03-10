package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando funcao init")
	n = 10
}

func main() {
	fmt.Println("Main Init")
	fmt.Println(n)
}
