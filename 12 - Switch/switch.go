package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Numero Ivalido"
	}
}

func main() {
	fmt.Println("SWITCH")

	dia := diaDaSemana(5)
	fmt.Println(dia)
}
