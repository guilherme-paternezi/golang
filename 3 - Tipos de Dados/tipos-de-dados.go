package main

import (
	"errors"
	"fmt"
)

func main() {

	//INT
	var numero1 int64 = -10000000000000
	fmt.Println(numero1)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Float
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1234.5678
	fmt.Println(numeroReal3)

	//STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "T"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	//BOOLEAN
	var booleano1 bool // por padrão o boolenao vem FALSE
	fmt.Println(booleano1)

	//ERROR
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
