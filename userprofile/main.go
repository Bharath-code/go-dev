package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Likes    int
	Comments []string
	Friends  []int
}

func main() {
	start := time.Now()
	userProfile, err := fetchUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userProfile)
	fmt.Println("fetching user profile took....", time.Since(start))
}

type Response struct {
	data any
	err  error
}

func fetchUserProfile(id int) (*UserProfile, error) {
	var (
		respch = make(chan Response, 3) // 3 because three function call comments, likes, friends async-lly
		wg     = &sync.WaitGroup{}
	)

	// we are doing  3 request inside their own goroutine
	go getComments(id, respch, wg)
	go getFriends(id, respch, wg)
	go getLikes(id, respch, wg)

	//Adding 3 to the waitgroup i.e we calling above three functions
	wg.Add(3)
	wg.Wait()     // block until wg counter == 0 then unblock , i.e wg.Done() happens in the functions bellow it will decrement wg.Add(3) one by one
	close(respch) // we should close the channel because if we don't then channel excepts more message and leads to deadlock
	userProfile := &UserProfile{}
	for resp := range respch {
		// type conversion
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}
	}

	return userProfile, nil
}

func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	likes := 33
	respch <- Response{
		data: likes,
		err:  nil,
	}
	wg.Done()
}

func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comment := []string{
		"Hello guys",
		"I'm fine , how do you do..",
		"I'm alright",
	}
	respch <- Response{
		data: comment,
		err:  nil,
	}
	wg.Done()
}

func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	friends := []int{34, 89, 432, 190}

	respch <- Response{
		data: friends,
		err:  nil,
	}
	wg.Done()
}
