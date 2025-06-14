package main

import "fmt"

func main() {
	//tipos de variaveis implicito ou explicito
	var variavel1 string = "conteudo que esta dentro da variavel 1"
	variavel2 := "variavel2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "teste 1"
		variavel4 string = "Teste 2"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "eu sou a variavel 5", "eu sou a varivel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
