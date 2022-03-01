package main

import "fmt"

func main() {
	a := 4
	b := squareVal(a)

	fmt.Println(a)
	fmt.Println(b)
}

func squareVal(v int) int {
	v *= v
	fmt.Println(v)
	return v
}
