package main

import "fmt"

var x = [5]int{1, 2, 3, 4, 5}

func array() {
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
