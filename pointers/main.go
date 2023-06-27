package main

import "fmt"

/*
	type CustomMap struct {
		data map[string]int
	}

Here string and int are hardcoded to make to it scalable and maintable it need to use
generic to accomadate any type
*/

/*
CustomMap[K comparable, V any] -> this is generics
comaparable is hashmap to get the hashed key value to compare
*/

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (c *CustomMap[K, V]) Insert(k K, v V) error {
	c.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func foo[T int](val T) {
	fmt.Println(val)
}

func main() {
	s := NewCustomMap[string, int]()
	s.Insert("fi", 3)
	s.Insert("hi", 10)

	d := NewCustomMap[int, float64]()
	d.Insert(23, 78.432102)
	d.Insert(19, 18.4)

	foo(1)
}
