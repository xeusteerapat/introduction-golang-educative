package main

import "fmt"

func main() {
	cities := []string{}
	cities = append(cities, "Bangkok", "Chiang Mai", "Nan") // append one of more
	fmt.Println(cities)                                     // [Bangkok Chiang Mai Nan]

	// And you can also append a slice to another using an ellipsis
	musics := []string{"Metal", "Disco"}   // slice of string
	otherMusic := []string{"Funk", "Rock"} // slice of string
	musics = append(musics, otherMusic...)

	fmt.Printf("%q\n", musics) // ["Metal" "Disco" "Funk" "Rock"]

	// Length of slice
	fmt.Println(len(musics))
}
