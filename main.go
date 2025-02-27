package main

import "fmt"

// Main entrance
func main() {
	fmt.Println("---------- ChatOyO Chat Server V0.3 ----------")
	server := NewServer("127.0.0.1", 2121) // Read from config file
	server.Run()
}
