package main

import "fmt"

func main() {
	x := 10
	square(&x)
	fmt.Println(x)
}

func square(x *int) *int {
	z := (*x) * (*x)
	return &z
}
