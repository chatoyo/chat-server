package main

import "fmt"

// Main entrance
func main() {
	fmt.Println("---------- ChatOyO Chat Server V0.3 ----------")
	config := LoadConfig()

	server := NewServer(config) // Read from config file

	server.Run()
}
