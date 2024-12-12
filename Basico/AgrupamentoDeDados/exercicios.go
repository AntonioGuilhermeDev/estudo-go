package main

import "fmt"

func exercicio1() {
	fmt.Println("Exercicio 1 \n")
	meuArray := [5]int{10, 20, 30, 40, 50}
	for index, value := range meuArray {
		fmt.Print(index, value, "\n")
	}
	fmt.Printf("O tipo do array é %T \n\n", meuArray)
}

func exercicio2() {
	fmt.Println("Exercicio 2 \n")
	meuSlice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for index, value := range meuSlice {
		fmt.Print(index, value, "\n")
	}
	fmt.Printf("\nO tipo do slice é %T \n\n", meuSlice)

	fmt.Println("Exercicio 3 \n")

	fatia1 := meuSlice[:3]
	fatia2 := meuSlice[4:]
	fatia3 := meuSlice[1:7]
	fatia4 := meuSlice[2:9]
	fatia5 := meuSlice[2 : len(meuSlice)-1]

	fmt.Println(fatia1)
	fmt.Println(fatia2)
	fmt.Println(fatia3)
	fmt.Println(fatia4)
	fmt.Println(fatia5)
	fmt.Println("")

}

func exercicio4() {
	fmt.Println("Exercicio 4 \n")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := []int{56, 57, 58, 59, 60}

	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("")
}

func exercicio5() {
	fmt.Println("Exercicio 5 \n")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	xSlice1 := x[:3]
	fmt.Println(xSlice1)
	xSlice2 := x[6:]
	fmt.Println(xSlice2)
	y := append(xSlice1, xSlice2...)
	fmt.Println(y)
	fmt.Println("")
}

func exercicio6() {
	fmt.Println("Exercicio 6 \n")

	x := make([]string, 26, 26)
	x = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceara", "Espirito Santo", "Goias", "Maranhao", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Para", "Paraiba", "Parana", "Pernambuco", "Piaui", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondonia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}
	fmt.Printf("Tamanho: %d, \n Capacidade: %d \n\n", len(x), cap(x))

	for i := 1; i < len(x); i++ {
		fmt.Println(x[i])
	}

	fmt.Println("")

}

func exercicio7() {
	fmt.Println("Exercicio 7 \n")

	x := [][]string{
		[]string{"Antonio", "Guilherme", "Música"},
		[]string{"Joao", "Lima", "Skate"},
		[]string{"Ana", "Sousa", "Games"},
	}

	for _, v := range x {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}

	fmt.Println("")
}

func exercicio8() {
	fmt.Println("Exercicio 8 \n")

	x := map[string][]string{
		"silva_joao": []string{
			"Literatura",
			"Musica",
			"Idiomas",
		},
		"lima_antonio": []string{
			"Idiomas",
			"Programação",
			"Música",
			"Jogos",
		},
	}

	for k, v := range x {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}
