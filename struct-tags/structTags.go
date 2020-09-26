package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// User struct takes name, password and preferredfish fields to create a user profile
type User struct {
	name          string
	password      string
	preferredFish []string
	createdAt     time.Time
}

func main() {
	user1 := &User{
		name:      "Sammy the Shark",
		password:  "fishareawesome",
		createdAt: time.Now(),
	}

	output, err := json.MarshalIndent(user1, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
