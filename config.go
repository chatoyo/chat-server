package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server struct {
		Ip                string `json:"ip"`
		Port              int    `json:"port"`
		MaxPendingSeconds int64  `json:"maxPendingSeconds"`
	}
}

func LoadConfig() (config *Config) {
	data, err := os.ReadFile("./config.json")

	if err != nil {
		fmt.Println("Error opening config file:", err)
		return nil
	}

	err = json.Unmarshal(data, &config)
	printServerConfig(config)

	return config
}

func printServerConfig(config *Config) {
	fmt.Println("[INFO] Loaded config file.")
	fmt.Println("[INFO] Server address:", config.Server.Ip)
	fmt.Println("[INFO] Server port:", config.Server.Port)
}
