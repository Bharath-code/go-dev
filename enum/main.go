/*
	learn enumeration
*/

package main

import "fmt"

//	weapon type
//	axe
//	sword
// 	wooden stick
// 	knife

type WeaponType int

const (
	Axe WeaponType = iota // basically it increment starts from 0 to number variables declared, instead of assigning each variables with value in numeric
	Sword
	WoodenStick
	Knife
)

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 1
	case Knife:
		return 60
	default:
		panic("weapon doesn't exist")
	}
}

func main() {

	fmt.Printf("damage of weapon (%d) (%d):\n", Axe, getDamage(Axe))
	fmt.Printf("damage of weapon (%d) (%d):\n", Sword, getDamage(Sword))
	fmt.Printf("damage of weapon (%d) (%d):\n", WoodenStick, getDamage(WoodenStick))

}
