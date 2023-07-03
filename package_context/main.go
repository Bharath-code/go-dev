package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "user", "bharath") // pass value to the context key value pair , like passing objects to api call in js
	result, error := fetchUserId(ctx)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("the response took %v -> %+v \n", time.Since(start), result)

}

func thirdpartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 90)
	return "user id 1", nil
}

func fetchUserId(ctx1 context.Context) (string, error) {
	/*
		----Context is very important----0
		1. while calling an api service which a third party , we may not know when it will give response because we don't have
		control over it.
		2. Ideal UX for response time is 100ms, so we can use context here to control the api call and improve the UX.
		3. context.Background() is parent context if we dont have any
		4. defer cancel() is a best practice because other function may depend on this context , if that call takes
		forever our whole experience go for a toss. So it cancel the call at the end
	*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	val := ctx1.Value("user")
	fmt.Println("the value is : ", val)

	type result struct {
		userId string
		err    error
	}

	resultch := make(chan result, 1)

	// third party api should be in async mode like used in js (async await)
	// since it's non blocking we need channel to communicate
	go func() {
		res, err := thirdpartyHTTPCall()
		resultch <- result{
			userId: res,
			err:    err,
		}
	}()

	select {
	// Done()
	// 1. -> context timeout is exceeded
	// 2. -> the context has been manually cancelled -> cancel()
	case <-ctx.Done():
		return "", ctx.Err() // this error is from line 35
	case res := <-resultch:
		return res.userId, res.err // this error is from line 50 , http call is done successfully but there is an error

	}

}
