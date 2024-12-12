package main

import "fmt"

func exercicio1() {
	x := 10

	fmt.Printf("%d, %b, %x", x, x, x)
}

func exercicio2() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exercicio3() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%U\n", i)
		}
	}
}

func exercicio4() {
	for i := 2006; i <= 2024; i++ {
		fmt.Println(i)
	}
}

func exercicio5() {
	birthday := 2006
	today := 2024
	for {
		if birthday == today {
			break
		}
		fmt.Println(birthday)
		birthday++
	}
}

func exercicio6() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i, "\t", i%4)
	}
}

func exercicio7() {
	hora := 15
	if hora < 24 {
		if hora >= 0 && hora <= 5 {
			fmt.Println("Boa madrugada")
		} else if hora >= 6 && hora <= 12 {
			fmt.Println("Bom Dia!")
		} else if hora >= 13 && hora <= 18 {
			fmt.Println("Boa Tarde!")
		} else {
			fmt.Println("Boa Noite!")
		}
	} else {
		fmt.Println("O dia só tem 24 horas")
	}
}

func exercicio8() {
	hora := 50

	switch {

	case hora >= 0 && hora <= 5:
		{
			fmt.Println("Boa madrugada")
		}
	case hora >= 6 && hora <= 12:
		{
			fmt.Println("Bom Dia!")
		}
	case hora >= 13 && hora <= 18:
		{
			fmt.Println("Boa Tarde!")
		}
	case hora >= 19 && hora <= 23:
		{
			fmt.Println("Boa Noite!")
		}
	default:
		{
			fmt.Println("O dia só tem 24 horas")
		}
	}
}

func exercicio9() {
	esporteFavorito := "Futebol"

	switch esporteFavorito {
	case "Futebol":
		{
			fmt.Println("Futebol")
		}
	case "Basquete":
		{
			fmt.Println("Basquete")
		}
	case "Volei":
		{
			fmt.Println("Volei")
		}
	case "Tenis":
		{
			fmt.Println("Tenis")
		}
	case "Ping Pong":
		{
			fmt.Println("Ping Pong")
		}
	default:
		{
			fmt.Println("Esporte não reconhecido")
		}
	}
}

func exercicio10() {
	n := 50

	for i := 1; i < n; i++ {
		fmt.Printf("%d + %d = %d\n ", i, i+1, i+(i+1))
	}
}

func exercicio11() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d X %d = %d\n ", i, j, i*j)
		}
		fmt.Println("")
	}
}
