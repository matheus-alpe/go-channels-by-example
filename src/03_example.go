package examples

import (
	"fmt"
)

type Server struct {
	users  map[string]string
	userch chan string
	quitch chan struct{}
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
		quitch: make(chan struct{}),
	}
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {
loop:
	for {
		select {
		case user := <-s.userch:
			fmt.Println("add user:", user)
			s.addUser(user)
		case <-s.quitch:
			fmt.Println("stopping")
			break loop
		}
	}
}

func (s *Server) addUser(user string) {
	s.users[user] = user
}

func ExecuteExample03() {
	fmt.Println("example 03")

	server := NewServer()
	server.Start()

	close(server.quitch)

	select {}
}
