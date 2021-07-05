package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(mySlice) // [1 2 3 4 5 6]

	fmt.Println(mySlice[1:4]) // [2 3 4] from index 1 to exclude index 4

	fmt.Println(mySlice[:3]) // [1 2 3] from beginning to exclude index 3

	fmt.Println(mySlice[4:]) // [5 6] from index 4 to the end

	name := [4]string{
		"Anthony Kiedis",
		"Flea",
		"Josh Klinghoffer",
		"Chad Smith",
	}

	fmt.Println(name) // [Anthony Kiedis Flea Josh Klinghoffer Chad Smith]

	a := name[0:2]    // slice a
	b := name[1:3]    // slice b
	fmt.Println(a, b) // [Anthony Kiedis Flea] [Flea Josh Klinghoffer]

	// Change value by specific index
	b[1] = "John Frusciante"
	fmt.Println(a, b) // [Anthony Kiedis Flea] [Flea John Frusciante]
	fmt.Println(name) // [Anthony Kiedis Flea John Frusciante Chad Smith] *note* name also change

	cities := make([]string, 3) // use make keyword to create empty slice

	cities[0] = "Santa Monica"
	cities[1] = "Venice"
	cities[2] = "Los Angeles"

	fmt.Printf("%q", cities)
}
