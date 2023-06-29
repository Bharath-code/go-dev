/*
  - go routines
    teaches the async properties in golang -> concurrency

like ->
1. fetch https resources
2. fetch db resources

  - channels
    since go rountines are non blocking we can't save those values returned
    because it's async , so channels comes in -> it's like tunnel of passage
    of values that can be given at any point of time and it has two types
    1. unbuffered channel -> it is blocking , tunnel is full -> this has no space send values without a box
    2. buffered channel -> it giving size to the variable -> store the item in a box and send across the tunnel
*/
package main

import (
	"fmt"
	"time"
)

func fetchResources(id int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", id)
}

func main() {
	//go fetchResources(1) -> async like it in JS async await
	/* go func(){
		fetchResources(1)
	}() -> can also written like this both methods are same -> it like IFFE in JS*/

	//variable that stores the channel it's good practice to end with "ch"
	// we can pass any type to channel
	// channel will always block when its full ,when its full below programs won't execute

	// buffered channel -> returns address like 0x14000064180 to return value use <- while assigning to local var
	resultch1 := make(chan string, 10)
	resultch1 <- "bar"
	result := <-resultch1
	fmt.Println("buffered channel", result)

	resultch := make(chan string) // unbuffered channel

	go func() {
		result := <-resultch
		fmt.Println("unbuffered channel", result)
	}()
	resultch <- "foo"

}
