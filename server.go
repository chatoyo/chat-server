package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	Message chan string
}

// Factory Constructor
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (server *Server) ListenMsg() {
	for {
		msg := <-server.Message

		server.mapLock.Lock()
		for _, client := range server.OnlineMap {
			client.C <- msg
		}

		server.mapLock.Unlock()
	}
}

func (server *Server) QueueBroadcastMsg(user *User, message string) {
	sendMsg := "[" + user.conn.RemoteAddr().String() + "] " + user.Name + " : " + message
	server.Message <- sendMsg
}

func (server *Server) Handler(conn net.Conn) {
	fmt.Printf("New connection from %s\n", conn.RemoteAddr())

	user := NewUser(conn, conn.RemoteAddr().String())

	// User Login, add to map
	server.mapLock.Lock()
	server.OnlineMap[user.Name] = user
	server.mapLock.Unlock()

	//Broadcasting the login msg
	server.QueueBroadcastMsg(user, "Login")

	// Block
	select {}
}

// Run a new Server
func (server *Server) Run() {
	listener, err := net.Listen("tcp", server.Ip+":"+strconv.Itoa(server.Port))
	if err != nil {
		fmt.Printf("net.Listen error: %v\n", err)
		return
	}

	defer listener.Close()

	go server.ListenMsg()

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
