/*
When to use pointers
How to use pointers
why to use pointers
*/
package main

import "fmt"

type Player struct {
	HP int
}

// pointers occupy 8 bytes as compared to call by value it's size increases and performance decreases(because it copies and used)
func TakeDamage(player *Player, amount int) error {
	player.HP -= amount
	fmt.Println("New HP -> ", player.HP) // -> before pointer -> New HP it would be 90
	return nil
}

/*Part - 2

Debug the pointers
how to handle pointers
*/

type Database struct {
	user string
}
type Server struct {
	db *Database //uintptr -> 8 byte long
}

func (s *Server) GetUserfromDB() string {
	// golang is going to dereference the db pointer
	//it's going to lookup the memory address of the pointer
	//it's always good practice to check if it's nil and proceed
	if s.db == nil {
		//panic("db is connection is not intialized") -> it's tell us what's the error
		return ""
	}
	return s.db.user
}

func main() {
	// we can use '&' in Player struct(player := &Player) or while passing the player variable
	player := Player{
		HP: 100,
	}
	//here we are sending only copy of the player, it will give the desired result see call by value and call by reference
	TakeDamage(&player, 10)
	fmt.Printf("%+v\n", player) //before pointer -> New HP it would be 100
	s := &Server{
		//db: &Database{}, -> this should be done to avoid error
	}
	s.GetUserfromDB()
}
