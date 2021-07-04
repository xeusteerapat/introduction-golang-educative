package main

import "fmt"

func main() {
	var (
		name     string = "Teerapat Prommarak"
		age      int    = 35
		location string = "Bangkok"
	)

	// Or implicit type assignment with :=
	name2, location2 := "Teerapat Prommarak", "Bangkok"
	age2 := 35
	fmt.Printf("%s age %d from %s ", name2, age2, location2)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(location)

	// variables can contain any type includes function
	myFunc := func() {
		fmt.Println("From my func")
	}

	myFunc()

	name6 := "Caprica-Six"
	aka := fmt.Sprintf("Number %d", 6)
	fmt.Printf("%s is also known as %s",
		name6, aka)
}
