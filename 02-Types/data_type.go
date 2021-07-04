package main

import (
	"fmt"
	"math/cmplx"
)

var (
	goIsFun bool       = true
	maxInt  uint64     = 1<<64 - 1
	complex complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n" // print type(%T) and value(%v)

	fmt.Printf(f, goIsFun, goIsFun) // bool(true)
	fmt.Printf(f, maxInt, maxInt)   // uint64(18446744073709551615)
	fmt.Printf(f, complex, complex) // complex128((2+3i))

	// Data type conversion
	i := 42
	j := float64(i)
	k := uint(j)

	fmt.Printf("i %T : %v\n", i, i)
	fmt.Printf("j %T : %v\n", j, j)
	fmt.Printf("k %T : %v\n", k, k)
}
