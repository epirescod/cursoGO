package main

import "fmt"

func funcao1() {
	fmt.Println("executando a função 1")
}

func funcao2() {
	fmt.Println("executando função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculado o resultado sera alterado")
	fmt.Println("entrando na função para verificar se o alunon esta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {

		return true
	}

	return false
}

func main() {
	defer funcao1()
	// defer dignifica adiar a execução desta função até o ultimo momento possivel = ADIAR
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
