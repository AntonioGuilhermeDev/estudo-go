package main

import "fmt"

func funDefer() {
	defer fmt.Println("1")
	defer fmt.Println("4")
	fmt.Println("2")
	fmt.Println("3")
}
