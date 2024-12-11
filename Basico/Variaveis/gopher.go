package main

import "fmt"

var f = 10

func gopher() {

	x := "ola" // --> GOpher: declaração da variavel, funciona apenas em codeblocks
	y := 20
	z := false

	fmt.Printf("x = %v, %T\n", x, x)
	fmt.Printf("y = %v, %T\n", y, y)
	fmt.Printf("z = %v, %T\n", z, z)

	x = "ola mundo" //--> Atribuição
	fmt.Printf("x = %v, %T\n", x, x)

	y, a := 20, true // Pode ser usado como atribuição se declarar pelo menos uma variavel nova junto
	fmt.Printf("y = %v, %T\n", y, y)
	fmt.Printf("z = %v, %T\n", a, a)

	fmt.Printf("f = %v, %T\n", f, f)
}
