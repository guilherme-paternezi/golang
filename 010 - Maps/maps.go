package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	usuario := map[string]string{
		"nome":      "Guilherme",
		"sobrenome": "Paternezi",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Guilherme",
			"segundo":  "Comicio",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 3",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")

}
