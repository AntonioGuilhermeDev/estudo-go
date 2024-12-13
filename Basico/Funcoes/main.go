package main

import "fmt"

func main() {
	total, quantos, oi := soma("tarde", 10, 10, 1, 2, 3, 5)
	fmt.Println(total, quantos, oi)
	funDefer()
}
