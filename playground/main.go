package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(i, j)

	p := &i // set the pointer points to i variable

	fmt.Println(&p)       // 0xc0000ae008
	fmt.Println(*p)       // return the value that the pointer is pointing to
	fmt.Printf("%T\n", p) // *int -> int based pointer type
	*p = 21
	fmt.Println(i) // i = 21
}
