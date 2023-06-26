package main

import "fmt"

type Position struct {
	x, y int
}

type Entity struct {
	name    string
	id      string
	version string
	Position
}

type SpecialEntity struct {
	specialField float64
	Entity
}

func main() {
	e := SpecialEntity{
		specialField: 34.89,
		Entity: Entity{
			name:    "jio",
			id:      "id 1",
			version: "ver 1.1",
			Position: Position{
				x: 150,
				y: 300,
			},
		},
	}
	e.version = "ver 2.0"
	fmt.Printf("%+v\n", e.version)
}
