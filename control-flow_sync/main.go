package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // because struct{} is zero bytec, bool is 1 or 2 bytes
	msgch  chan string
}

// start a new instance of the server like a constructor gives an object
func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) start() error {
	fmt.Println("Server is starting.....")
	s.loop()
	return nil
}

func (s *Server) quit() error {
	close(s.quitch)
	//or s.quitch <- struct{}{}
	return nil
}

func (s *Server) loop() error {
mainloop: // if we break from the select the main for will run infinitely , so we mainloop(name anything) and break fully from the for and select
	for {
		select {
		// why we don't assign the channel here for quitch is, it is a struct just a signal
		case <-s.quitch:
			fmt.Println("quitting server")
			break mainloop
		case msg := <-s.msgch:
			fmt.Println("in msg processing")
			s.handleMessage(msg)
		}
	}
	return nil
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we received the message ", msg)
}

func main() {
	/*
		1. creates a new instances of a server
		2.  starts the server
		3. send msg hello world
		4. delay server for 5 secs
		5. quit the server
	*/
	serve := newServer()
	go func() {
		serve.sendMessage("hello world")
		time.Sleep(time.Second * 5)
		serve.quit()
	}()
	// starts sync and starts before the above go func as latter is non blocking async function
	serve.start()

}
