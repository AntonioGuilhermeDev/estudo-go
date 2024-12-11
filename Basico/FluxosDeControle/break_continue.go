package main

import "fmt"

func forBreak() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			fmt.Println("BREAK! \n")
			break
		}
	}
}

func forContinue() {
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
