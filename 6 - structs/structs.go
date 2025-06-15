package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("arquivo structs")

	var u usuario
	u.nome = "Emanuel Pires"
	u.idade = 25
	fmt.Println(u)

	u2 := usuario{"Emanuel Pires", 25}
	fmt.Println(u2)

	u3 := usuario{nome: "Emanuel Pires"}
	fmt.Println(u3)

}
