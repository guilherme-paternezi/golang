package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Guilherme"
	u.idade = 20
	fmt.Println(u)

	enderecoU2 := endereco{"Rua Rio de Janeiro", 219}

	usuario2 := usuario{"Guilherme", 20, enderecoU2}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Paterezi"}
	fmt.Println(usuario3)
}
