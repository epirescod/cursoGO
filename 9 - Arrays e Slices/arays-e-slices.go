package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"p1", "p2", "p3", "p4", "p5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18, 19, 30) // slice é uma fatia ou pedaço de um array

	fmt.Println(slice)

	slice2 := array2[1:4]
	fmt.Println(slice2)

	//Array internos, aprendendo sobre utilizar o MAKE

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5) // adicionando mais um valor no slice3
	slice3 = append(slice3, 6) // adicionando mais um valor no slice3 tornando a capacidade dele pra 12

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length ou tamanho
	fmt.Println(cap(slice3)) // capacidade do slice3

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
