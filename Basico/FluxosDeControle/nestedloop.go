package main

import "fmt"

func nestedLoop() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora:", horas, "\n")
		fmt.Println("Minutos: \n")
		for minutos := 0; minutos < 60; minutos++ {
			fmt.Println(minutos, "\n")
		}
		fmt.Println("\n")
	}

	for meses := 1; meses <= 12; meses++ {
		fmt.Println("Mes:", meses, "\n")
		fmt.Println("Dias:", "\n")
		for dias := 1; dias <= 30; dias++ {
			fmt.Println(dias, "\n")
		}
		fmt.Println("\n")
	}
}
