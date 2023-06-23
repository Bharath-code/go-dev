/*
	learn interface
	structs
	interface
	attaching function to structs
	access associated methods
	power of interface

*/

/*
Struct: A struct is like a blueprint for an object. It defines the fields that the object will have, and the types of those fields. For example, a struct for a person might have fields for the person's name, age, and address.
Interface: An interface is like a contract. It defines the methods that an object must have in order to be considered a member of that interface. For example, an interface for a shape might define methods for getting the shape's area and perimeter.
Here are some more concrete analogies:

A struct is like a car. It has a specific set of parts (fields), and those parts have specific types (the types of the fields). A car can also be used to perform certain actions (methods).
An interface is like a driver's license. It doesn't specify what kind of car you can drive, but it does specify that you have the skills and knowledge to drive a car safely. In Golang, an interface specifies that an object has a certain set of methods, but it doesn't specify what those methods do.
*/
package main

import "fmt"

// interface -> it's like a contract
type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

// types of db store structs
type MongoDBNumberStore struct {
}

type PostgresNumberStore struct {
	// postgres values (db connection)
}

// attaching the respective function to their structs
func (p PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5, 6}, nil
}

func (p PostgresNumberStore) Put(number int) error {
	fmt.Println("store the number into the postgres storage")
	return nil
}

// this method will get attached to the MongoDBNumberStore struct
func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into the mongoDB storage")
	return nil
}

type ApiServer struct {
	numberStorer NumberStorer
}

func main() {
	apiServer := ApiServer{
		// plug and play the db store you wanted , no need to change everything and hardcode it and it's simple
		numberStorer: PostgresNumberStore{},
	}
	//why err is used twice but doesn't throw error
	numbers, err := apiServer.numberStorer.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(numbers)
	if err := apiServer.numberStorer.Put(1); err != nil {
		panic(err)
	}

}
