package main

import "fmt"

type hotdog int // --> Tipo customizado
var v hotdog

func tipos() {
	// é impossivel mudar tipos de variaveis
	// Tipagem estática
	// Tipos primitivos: int, string, bool
	// Tipos Compostos: Slice, array, Struct, Map

	w := "hi"
	x := 10   // --> Integer
	y := 20.5 // --> Float
	z := true // --> Bool

	fmt.Println(v, w, x, y, z)
	fmt.Printf("v %v is of type %T\n", v, v)
	fmt.Printf("w %v is of type %T\n", w, w)
	fmt.Printf("x %v is of type %T\n", x, x)
	fmt.Printf("y %v is of type %T\n", y, y)
	fmt.Printf("z %v is of type %T\n", z, z)

	/*

		Declaração --> Alocamento do valor na memoria
		Incialização --> Primeiro valor da variavel
		Atribuição --> Valor colocado depois da Inicialização

	*/

	/*
		valor zero

		Valor Presente na variavel por padrão antes de ser inicializada


		int: 0
		floats: 0.0
		booleans: false
		strings: ""
		pointers, functions, interfaces, slices, channels, maps: nil
	*/
}
