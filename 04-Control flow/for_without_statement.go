package main

import "fmt"

func main() {
	sum := 1

	for sum < 1000 { // iterate as long as sum<1000
		sum += sum
	}

	fmt.Println(sum)
}
