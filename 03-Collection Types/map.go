package main

import "fmt"

func main() {
	rhcp := map[string]int{
		"Anthony Kiedis":  58,
		"Flea":            58,
		"AChad Smith":     59,
		"John Frusciante": 51,
	}

	fmt.Printf("%#v\n", rhcp)

	for key, value := range rhcp {

		fmt.Printf("%s is %d years old\n", key, value)
	}
}
