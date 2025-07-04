package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A media é exatamente 6!") // mata a execução do programa
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pos execção")
}
