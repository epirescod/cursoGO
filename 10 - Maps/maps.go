package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Emanuel",
		"sobrenome": "Pires",
	}

	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Emanuel",
			"segundo":  "Pires",
		},

		"curso": {
			"nome":   "Engenharia da computação",
			"campus": "Uninove Vila Maria",
		},
	}

	fmt.Println(usuario2)

	// se você quiser deletar uma chave dentro de um map use a função abaixo

	delete(usuario2, "nome")
	fmt.Println(usuario2)
}
