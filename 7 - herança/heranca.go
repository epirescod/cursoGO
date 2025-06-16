package main

import "fmt"

//lembrando que o GO não tem heranças por não ser uma linguagem orientada a objeto

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("herança")

	p1 := pessoa{"João", "Pedro", 20, 170}
	fmt.Println(p1)

	e1 := estudante{p1, "Engeheiro", "USP"}
	fmt.Println(e1.altura)

}
