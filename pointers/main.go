package main

import "fmt"

type Player struct {
	HP int
}

func TakeDamage(player *Player, amount int) error {
	player.HP -= amount
	fmt.Println("New HP -> ", player.HP) // -> before pointer -> New HP it would be 90
	return nil
}

func main() {
	// we can use '&' in Player struct(player := &Player) or while passing the player variable
	player := Player{
		HP: 100,
	}
	//here we are sending only copy of the player, it will give the desired result see call by value and call by reference
	TakeDamage(&player, 10)
	fmt.Printf("%+v\n", player) //before pointer -> New HP it would be 100
}
