package main

import "fmt"

func slice() {

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 348756
	fmt.Println(slice[3])
}

func sliceForRange() {
	slice := []string{"banana", "maca", "uva"}

	for indice, valor := range slice {
		fmt.Println("No indice", indice, "temos o valor", valor)
	}
	slice = append(slice, "melancia")

	for indice, valor := range slice {
		fmt.Println("No indice", indice, "temos o valor", valor)
	}

	numSlice := []int{20, 25, 32, 40, 50}
	total := 0

	for _, valor := range numSlice {
		total += valor
	}
	fmt.Println("O valor total é", total)
}

func sliceASlice() {
	sabores := []string{"Peperonni", "Quatro Queijos", "Margerita", "Mozzarela", "Abacaxi"}

	fatia := sabores[0:3]

	fmt.Println(fatia)

	//Listando todos os elementos
	fatia = sabores[:]

	fmt.Println(fatia)

	// Listando todos os elementos com for
	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
	fmt.Println(fatia)

	//Removendo elemento no meio do slice
	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)
}

func appendSlice() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{7, 8, 9, 10}

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1 = append(slice1, 6)
	fmt.Println(slice1)

	// ...(unfurl ou enumeration) se refere aos elementos do slice, sem eles da erro
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
}

func makeSlice() {
	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))
}
func sliceMultiDimensional() {
	ss := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	//Lista 1
	fmt.Println(ss[1])

	// Lista 1 posição 1
	fmt.Println(ss[1][1])
}

func arraySubjacente() {

	primeiroSlice := []int{2, 3, 5, 7, 11, 13}
	segundoSlice := append(primeiroSlice[:2], primeiroSlice[4:]...)

	fmt.Println(segundoSlice)
	fmt.Println(primeiroSlice)
}

func maps() {

	pessoa := map[string]int{
		"joao":      1234567,
		"maria":     412578,
		"antonio":   484563,
		"gustavo":   848796,
		"nicolas":   548516,
		"guilherme": 456123,
	}

	fmt.Println(pessoa)
	fmt.Println(pessoa["joao"])

	pessoa["gopher"] = 412579

	for key, value := range pessoa {
		fmt.Println(key, value)
	}

	delete(pessoa, "guilherme")
	fmt.Println("Guilherme foi removido")

	if esta_presente, ok := pessoa["guilherme"]; !ok {
		fmt.Printf("Não está presente \n\n")
	} else {
		fmt.Println(esta_presente)
	}

	//False --> não presente
}
