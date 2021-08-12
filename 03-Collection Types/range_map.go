package main

import "fmt"

func main() {
	cities := map[string]int{
		"Bangkok":    10539000,
		"Chiang Mai": 127240,
		"Nan":        20212,
	}

	for key, value := range cities {
		fmt.Printf("%s has %d populations\n", key, value)
	}

	// 	Nan has 20212 populations
	// Bangkok has 10539000 populations
	// Chiang Mai has 127240 populations
}
