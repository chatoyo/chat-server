package main

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	Ip   string
	Port int
}

// Factory Constructor
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (server *Server) Handler(conn net.Conn) {
	// TODO: service
	fmt.Printf("New connection from %s\n", conn.RemoteAddr())
}

// Run a new Server
func (server *Server) Run() {
	listener, err := net.Listen("tcp", server.Ip+":"+strconv.Itoa(server.Port))
	if err != nil {
		fmt.Printf("net.Listen error: %v\n", err)
		return
	}

	defer listener.Close()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listener accept error: %v\n", err)
			continue
		}

		// do handler
		go server.Handler(conn)
	}

}
