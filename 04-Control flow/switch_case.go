package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	mins := now % 2

	switch mins {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	}

	printScore()
}

func printScore() {
	score := 7
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did you cheat?")
	default: //to be executed if no other case matches
		fmt.Println(score, " off the chart")
	}
}
