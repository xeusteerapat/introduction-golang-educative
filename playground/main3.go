package main

import "fmt"

func main() {
	a := 4
	fmt.Println(a)
	b := squareAdd(&a)
	fmt.Println(b)
}

func squareAdd(v *int) *int {
	*v *= *v
	fmt.Println(v)
	return v
}
