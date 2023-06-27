/*
	1. advanced interface
	2. typed function
*/

package main

import (
	"fmt"
)

type Putter interface {
	Put(id int, val any) error
	fmt.Stringer
}

type Storage interface {
	//embbed interface in another interface
	Putter
	Get(id int) (any, error)
}

type SimplePutter struct {
}

func (s SimplePutter) Put(id int, val any) error {
	return nil
}

func (SimplePutter) String() string {
	return ""
}

type FooStorage struct{}

func (s FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s FooStorage) Put(id int, val any) error {
	return nil
}
func (FooStorage) String() string {
	return ""
}

type BarStorage struct{}

func (s BarStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s BarStorage) Put(id int, val any) error {
	return nil
}
func (BarStorage) String() string {
	return ""
}

func updateValue(id int, val any, p Putter) error {
	return p.Put(id, val)
}

type Server struct {
	store Storage
}

func main() {
	s := Server{
		store: FooStorage{},
	}
	b := Server{
		store: BarStorage{},
	}
	sputter := SimplePutter{}
	updateValue(1, "foo", s.store)
	updateValue(2, "bar", b.store)
	updateValue(3, "foo", sputter)

}
