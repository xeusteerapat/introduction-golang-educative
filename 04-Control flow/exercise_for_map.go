package main

import (
	"fmt"
	"strconv"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	GetResult(users)
}

func GetResult([]string) string {
	//insert code here

	coinsForUser := func(name string) int {
		var total int
		for i := 0; i < len(name); i++ {
			switch string(name[i]) {
			case "a", "A":
				total++
			case "e", "E":
				total++
			case "i", "I":
				total = total + 2
			case "o", "O":
				total = total + 3
			case "u", "U":
				total = total + 4
			}
		}
		return total
	}
	for _, name := range users {
		v := coinsForUser(name)
		if v > 10 {
			v = 10
		}
		distribution[name] = v
		coins = coins - v
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
	return strconv.Itoa(distribution["Matthew"]) + " " + strconv.Itoa(distribution["Sarah"]) + " " + strconv.Itoa(distribution["Augustus"]) + " " + strconv.Itoa(distribution["Heidi"]) + " " + strconv.Itoa(distribution["Emilie"]) + " " + strconv.Itoa(distribution["Peter"]) + " " + strconv.Itoa(distribution["Giana"]) + " " + strconv.Itoa(distribution["Adriano"]) + " " + strconv.Itoa(distribution["Aaron"]) + " " + strconv.Itoa(distribution["Elizabeth"])
}
