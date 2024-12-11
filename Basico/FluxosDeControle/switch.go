package main

import (
	"fmt"
)

func switchCase() {

	x := 5

	switch {
	case x < 5:
		fmt.Printf("x is less than 5 \n")
	case x > 5:
		fmt.Printf("x is greater than 5 \n")
	default:
		fmt.Printf("x is equal to 5 \n")
	}

	pessoa := "Guilherme"

	switch pessoa {
	case "Guilherme", "Joao":
		fmt.Println("Equipe 1")
	case "Maria", "Pedro":
		fmt.Println("Equipe 2")
	default:
		fmt.Println("Equipe 3")
	}
}

/*
SWITCH
- Pode avaliar uma expressao
- default switch statement == true (bool)
- não há fail through por padrao
- default
- cases compostos
*/
