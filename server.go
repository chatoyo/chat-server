package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
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

type UserStatus int

const (
	LOGIN = iota
	LOGOUT
)

func (server *Server) updateUserStatus(user *User, status UserStatus) {
	server.mapLock.Lock()
	defer server.mapLock.Unlock()

	switch status {
	case LOGIN:
		server.OnlineMap[user.Name] = user
		break
	case LOGOUT:
		delete(server.OnlineMap, user.Name)
		break
	default:
	}
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
	sendMsg := "[" + user.conn.RemoteAddr().String() + "] " + user.Name + ": " + message
	server.Message <- sendMsg
}

func (server *Server) SendMsg(user *User, message string) {
	user.conn.Write([]byte(message))
}

func (server *Server) ParseMsg(user *User, msg string) {
	// Parse User Msg
	if msg == "who" {
		server.mapLock.Lock()
		var onlineMsg string
		for _, user := range server.OnlineMap {
			onlineMsg += "[" + user.conn.RemoteAddr().String() + "] " + user.Name + " " + "ONLINE\n"
		}
		server.SendMsg(user, onlineMsg)
		server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]

		_, nameAlreadyExists := server.OnlineMap[newName]
		if nameAlreadyExists {
			server.SendMsg(user, "[FAIL] Rename failed. Duplicated name.\n")
		} else {
			server.mapLock.Lock()
			delete(server.OnlineMap, user.Name)
			server.OnlineMap[newName] = user
			server.mapLock.Unlock()
			server.SendMsg(user, "[SUCCESS] Rename successfully.\n")
		}
		user.Name = newName

	} else {
		// Send Ordinary Msg
		server.QueueBroadcastMsg(user, msg)
	}
}

func (server *Server) Handler(conn net.Conn) {
	fmt.Printf("[INFO][LOGIN] New connection from %s\n", conn.RemoteAddr())

	user := NewUser(conn, conn.RemoteAddr().String())

	// User Login, add to map
	server.updateUserStatus(user, LOGIN)

	//Broadcasting the login msg
	server.QueueBroadcastMsg(user, "Login")

	// Activity Channel
	alive := make(chan bool)

	// Receive client msg
	go func() {
		buffer := make([]byte, 1024)

		for {
			n, err := conn.Read(buffer)
			if n == 0 {
				server.QueueBroadcastMsg(user, "Logout")
				server.updateUserStatus(user, LOGOUT)

				fmt.Printf("[INFO][LOGOUT] %s\n", conn.RemoteAddr())
				return
			}
			// Illegal operate
			if err != nil && err != io.EOF {
				fmt.Println("Conn read error:", err)
				return
			}

			// Extract user msg and remove "\n"
			// Windows n=n-2, Linux n maybe equals to n-1?
			msg := string(buffer[:n-2])

			server.ParseMsg(user, msg)

			alive <- true
		}
	}()

	for {
		select {
		case <-alive: // Must on the top of Timer Handler, due to the execute sequence
			// Alive
			// Reset Timer
		case <-time.After(time.Second * 10):
			// Timeout
			// Force offline
			server.SendMsg(user, "[TIMEOUT] Online timeout, force offline.\n")
			close(user.C)
			conn.Close()
			return
		}
	}

	// Meaningless Block
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
