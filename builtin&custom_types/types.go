/*
	1. General types
	2. Structs
	3. Maps
	4. slices
	5. arrys
	5. custom types
*/

package main

import "fmt"

//built-in data types

var (
	floatVar32 float32 = 0.1
	floatVar64 float64 = 0.122 // this is used widely in current scenario
	name       string  = "Next"
	intVar32   int32   = 1
	intVar64   int64   = 345543
	intVar     int     = -1
	uintVar    uint    = 1 // no negative numbers are allowed in uint because it's unassigned
	uint32Var  uint32  = 1
	uint64Var  uint64  = 10
	uint8Var   uint8   = 255 // only three digit upto 255 is allowed
	byteVar    byte    = 255 // only three digit upto 255 is allowed
	runVar     rune    = 's' // only single quotes is applicable , if it's double quotes then it's a string it will throw error
)

// structs
type Player struct {
	name        string
	health      int
	attackPower float64
}

// custom function
func getHealth(player Player) int {
	return player.health
}

// custom types
type Weapon string

func getWeapon(weapon Weapon) string {
	return string(weapon)
}

func main() {

	fmt.Println("Maps")
	//Maps -> it's like object in js and we can create map using two methods
	//users := map[string]int{}
	users := make(map[string]int)
	users["foo"] = 30
	users["bar"] = 25
	users["nat"] = 20

	age, ok := users["baz"]
	if !ok {
		fmt.Println("baz not exist in the map")
	} else {
		fmt.Println("baz exists in the map")
	}

	//delete key in map
	delete(users, "foo")
	fmt.Println(users)

	//loop over the map
	for k, v := range users {
		fmt.Printf("the key is %s and value is %d \n", k, v)
	}
	fmt.Println("-----------------------------")
	fmt.Println("Structs")
	player := Player{
		name:        "Ronak",
		health:      100,
		attackPower: 78.4,
	}
	// %v is print verbose
	// %+v is print verbose and key name of the structs
	fmt.Printf("Health : %d\n", getHealth(player))
	fmt.Printf("Age : %v\n", age)

	fmt.Println("-----------------------------")
	fmt.Println("Array and slices")

	//slices -> similar to array in other langs , here it can grow dynamically
	numberSlice := []int{1, 2}
	number2Slice := make([]int, 4)
	numberSlice = append(numberSlice, 3)
	number2Slice = append(number2Slice, 10)
	fmt.Println(numberSlice)
	fmt.Println(number2Slice)

	//array -> unlike slice array cannot grow dynamically and can't use append method
	numArray := [2]int{}
	numArray[0] = 23
	numArray[1] = 45
	fmt.Println(numArray)
	fmt.Println("-----------------------------")
}
