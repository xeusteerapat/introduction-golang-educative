package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for i := range pow {
		if i%2 == 0 {
			continue // skip
		}
		pow[i] = 1 << uint(i)
	}

	fmt.Println(pow) // [0 2 0 8 0 32 0 128 0 512]
}
