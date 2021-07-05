package main

import "fmt"

func main() {
	var a [2]string         //array of size 2
	a[0] = "Hello"          //Zero index of "a" has "Hello"
	a[1] = "World"          //1st index of "a" has "World"
	fmt.Println(a[0], a[1]) // will print Hello World
	fmt.Println(a)          // will print [Hello World]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
