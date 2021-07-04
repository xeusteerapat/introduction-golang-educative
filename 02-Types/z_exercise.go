package main

import "fmt"

// import "encoding/json"

type User struct {
	Id             int
	Name, Location string
}

func (u User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

type Player struct {
	u      User
	GameId int
}

func GreetingsForPlayer(p Player) string {
	//insert code

	return p.u.Greetings()
}

func main() {
	player := Player{}
	player.u.Id = 42
	player.u.Name = "Teerapat"
	player.u.Location = "BKK"
	fmt.Println(player.u.Greetings())
	fmt.Println(player.u)
}
