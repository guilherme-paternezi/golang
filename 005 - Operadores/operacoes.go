package main

import "fmt"

func main() {

	// ARITMETICOS
	soma := 1 + 2
	subtarcao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtarcao, divisao, multiplicacao, restoDaDivisao)

	//ATRIBUIÇÃO
	var variavel1 string = "String1"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(
		1 > 2,
		1 >= 2,
		1 == 2,
		1 <= 2,
		1 > 2,
		1 != 2)

	//OPERADORES LÓGICOS
	fmt.Println(".................")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!falso)
	fmt.Println(!verdadeiro)

	//OPERADORES UNÁRIOS
	fmt.Println("..................")
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20

	numero *= 3
	numero /= 10
	numero %= 3
}
