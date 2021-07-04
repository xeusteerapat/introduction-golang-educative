package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (user *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		user.Name, user.Location)
}

type Player struct {
	User
	GameId int
}

func main() {
	player := Player{}
	player.Id = 42
	player.Name = "Teerapat"
	player.Location = "BKK"
	fmt.Println(player.Greetings())
}
