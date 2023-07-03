/*
	-> atomic value are same as mutex , but we should in same functions like counter and boolean , because may overuse it
	A wait group and a mutex are both synchronization primitives in Golang that are used to control access to shared resources. However, they have different purposes and are used in different scenarios.

A wait group is used to wait for a set of goroutines to finish executing.
When a goroutine is created, it calls Add(1) on the wait group.
When the goroutine finishes executing, it calls Done() on the wait group. The Wait() method on the wait group blocks until all of the goroutines that have called Add(1) have also called Done().

A mutex is used to lock a shared resource so that only one goroutine can access it at a time.
When a goroutine needs to access a shared resource, it calls Lock() on the mutex.
When the goroutine is finished accessing the resource, it calls Unlock() on the mutex.
This ensures that no two goroutines can access the resource at the same time and that the resource is always in a consistent state.

In general, a "wait group is used to coordinate the execution of goroutines, while a mutex is used to protect shared resources from concurrent access".
*/
package main

import (
	"sync"
)

type State struct {
	mu    sync.Mutex // remove this to use atomic values
	count int        //count int 32 for using atomic for this example
}

// lock and unlock is one operation is done at a time and unlock to perform another
func (s *State) setState(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count = i
	//atomic.AddInt32(&s.count,int32(i))
}

func main() {

}
