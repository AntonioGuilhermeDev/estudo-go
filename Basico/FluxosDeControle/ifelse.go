package main

import "fmt"

func ifElse() {
	if x := 60; x > 100 {
		fmt.Println("x is greater than 100")
	} else if x > 50 {
		fmt.Println("x is greater than 50 but is less than 100")
	} else {
		fmt.Println("x is less than 100")
	}
}
