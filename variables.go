/*
	This lesson objectives
	1. global variables
	2. local variables
	3. constant variables
	4. variable declaration

*/

package main

import "fmt"

// global variables -> outside of functions starts with "var" keyword
var (
	name             = "Bharath kumar"
	firstName string = "Bharath"
	lastName  string
)

// constant variables -> whose values cannot be changed and it's starts with lowercase unlike other languages which is uppercase
const (
	initialVersion int = 1
	keyLen         int = 10
)

func main() {

	fmt.Println("hello")
	fmt.Println(name)
	fmt.Println(firstName)
	fmt.Println(lastName)

	//local variables ->  these can be initialized using ":=" and assigned datatype is based on value initialized
	version := 2            //infer to int type
	otherVersion := "Bar"   // infer to string type
	anotherVersion := 10.12 // infer to float32 or float 64
	fmt.Println(version)
	fmt.Println(otherVersion)
	fmt.Println(anotherVersion)

	// default value of int is "0" , string is ""(empty string)
	// using var inside function is not advised
	var lastVersion int
	lastVersion = 10
	fmt.Println(lastVersion)

	fmt.Println(initialVersion)
	fmt.Println(keyLen)
}
