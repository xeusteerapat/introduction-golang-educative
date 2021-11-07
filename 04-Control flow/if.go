package main

import (
	"fmt"
	"math"
)

func sqrt(num float64) string {
	if num < 0 {
		return sqrt(-num) + "i"
	}

	return fmt.Sprint(math.Sqrt(num))
}

func main() {
	sqrtOfTwo := sqrt(2)
	sqrtOfNegativeFour := sqrt(-4)

	fmt.Println(sqrtOfTwo, sqrtOfNegativeFour) // 1.4142135623730951 2i
}
