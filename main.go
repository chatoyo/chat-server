package main

// Main entrance
func main() {
	server := NewServer("127.0.0.1", 2121) // Read from config file
	server.Run()
}
