package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // because struct{} is zero byte
	msgch  chan string
}

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
mainloop:
	for {
		select {
		// why we don't assign the channel here for quitch is it is a struct just a signal
		case <-s.quitch:
			fmt.Println("quitting server")
			break mainloop
		case msg := <-s.msgch:
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
	serve := newServer()
	go func() {
		time.Sleep(time.Second * 5)
		serve.quit()
	}()

	serve.start()

}
