package main

import (
	"fmt"
	"math"
)

func pow(num, expo, limit float64) float64 {
	if value := math.Pow(num, expo); value < limit { // if value less than limit then return value
		return value
	}

	return limit
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
