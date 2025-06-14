package main

import (
	"errors"
	"fmt"
)

func main() {
	//tipo de dado numericos inteiro (int8, int16, int32, int64) int sozinho ele usa arquitetura da sua maquina

	var numero int16 = 100
	fmt.Println(numero)

	//uint numero sem sinal

	var numer2 uint32 = 10000
	fmt.Println(numer2)

	// alias apelido para referencia
	// INT 32 = RUNE

	var numero3 rune = 123456

	fmt.Println(numero3)

	// BYTE = UINT8
	var numer4 byte = 123
	fmt.Println(numer4)

	//numeros reais são float seguindo a mesma caracterisca do int32...

	var numeroReal1 float32 = 123.40
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.40
	fmt.Println(numeroReal3)

	// fim dos numero reais

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'c'
	fmt.Println(char)

	//FIM STRINGS

	var texto int16
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
