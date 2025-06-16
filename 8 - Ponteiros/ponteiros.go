package main

import "fmt"

func main() {
	fmt.Println("Pronteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++ // adiciona valor na varivel1 apenas nela, sem alterar o valor da variavel2

	fmt.Println(variavel1)

	// ponteiro é uma referencia de memória

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro)

	variavel3 = 500
	fmt.Println(variavel3, ponteiro)

}
