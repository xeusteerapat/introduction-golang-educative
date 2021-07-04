package main

import "fmt"

func main() {
	var yourAge *int // dereference

	var myAge int = 35
	fmt.Printf("myAge=%p myAge=%v\n", &myAge, myAge) // 0xc0000ae008

	yourAge = &myAge // get the pointer of myAge variable

	fmt.Println(yourAge) // 0xc0000ae008 this will print the same memory address with my age
}
