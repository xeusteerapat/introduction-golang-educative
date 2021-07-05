package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	// [] 0 0
	if z == nil {
		fmt.Println("nil!")
	}
	// nil!
}
