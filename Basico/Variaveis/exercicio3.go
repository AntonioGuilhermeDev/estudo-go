package main

import "fmt"

var a = 10
var b = "James Bond"
var c = true

func exercicio3() {
	s := fmt.Sprintf("%v, %v, %v", a, b, c)
	fmt.Println(s)
}
