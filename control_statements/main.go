/*
	learn control statement
	1. for loops
	2. range
	3. break
*/

package main

import "fmt"

func main() {

	//classic for statement
	fmt.Println("----classic for statement----------")
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("----for with range----------")
	// for with range
	names := []string{"a", "b", "c", "d"}
	//range will give index and value , so to omit the index we can use "_" so complier doesn't show warning like variable declared but not used
	for _, name := range names {
		fmt.Println(name)
	}
	fmt.Println("----for with if and break----------")
	// for with range
	diffNames := []string{"a", "b", "c", "d"}
	//range will give index and value , so to omit the index we can use "_" so complier doesn't show warning like variable declared but not used
	for _, name := range diffNames {
		if name == "a" {
			break
		}
	}

	fmt.Println("----switch statement----------")
	checkName := "Salas"
	// case statement has default break statement unlike js, otherwise it print everything
	switch checkName {
	case "Verak":
		fmt.Println("The name is ", checkName)
	case "Vinak":
		fmt.Println("The name is ", checkName)
	default:
		fmt.Println("the name is default => ", checkName)
	}
}
