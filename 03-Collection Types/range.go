package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow { //loops over the length of pow
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10) // [0 0 0 0 0 0 0 0 0 0]

	for i := range pow {
		pow[i] = 1 << uint(i) // [1 2 4 8 16 32 64 128 256 512]
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value) // print only value of pow
	}
}
