package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

// type Player with one additional attribute

type Player struct {
	User       // Composing User Struct
	GameId int // added attribute
}

func main() {
	p := Player{} // initializing
	p.Id = 42
	p.Name = "Teerapat"
	p.Location = "BKK"
	p.GameId = 471983
	fmt.Printf("%+v", p)
}
