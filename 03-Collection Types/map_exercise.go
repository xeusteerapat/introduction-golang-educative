package main

import (
	"fmt"
	"strings"
	//    "encoding/json"
	//    "strconv"
)

func WordCount(s string) map[string]int {
	// write your solution here
	words := strings.Fields(s)
	counter := map[string]int{}

	for _, word := range words {
		counter[word]++
	}

	return counter
}

func main() {
	ans := WordCount("My name is Teerapat")
	fmt.Println(ans)
}
