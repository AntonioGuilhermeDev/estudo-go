package main

import "fmt"

type damage int

var d damage
var e int

func exercicio4() {
	fmt.Printf("%v, %T\n", d, d)
	d = 42
	fmt.Printf("%v, %T\n", d, d)
	e = int(d)
	fmt.Printf("%v, %T\n", e, e)
}
