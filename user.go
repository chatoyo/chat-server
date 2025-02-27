package main

import "net"

type User struct {
	Name string
	Id   int
	C    chan string
	conn net.Conn
}

func NewUser(conn net.Conn, name string) *User {
	user := &User{
		Name: name,
		conn: conn,
		C:    make(chan string),
	}

	// run a routine to listen to client msg
	go user.ListenMsg()

	return user
}

// ListenMsg Listen to User Channel. Once new msg get, send to the client
func (user *User) ListenMsg() {
	for {
		msg := <-user.C
		user.conn.Write([]byte(msg + "\n"))
	}
}
