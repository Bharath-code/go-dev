package main

import (
	"fmt"
	"sync"
	"testing"
)

/*
1. race condition is at the same time , go rountine trying to write in the same variable
2. waitgroup -> since go rountine are async , we dont know what happened which will execute to make it sync with us we use this.
*/
func TestState(t *testing.T) {
	state := &State{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			state.setState(i)
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Printf("%+v \n", state)
}
