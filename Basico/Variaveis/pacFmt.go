package main

import "fmt"

func pacFmt() {
	x := "texto"
	y := "texto2"
	z := fmt.Sprint(x, " ", y)

	fmt.Println(x)
	fmt.Println(z)

}
