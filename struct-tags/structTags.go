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

// UserWithAnnotations is a carbon copy of User but with struct annotations to illustrate their use
type UserWithAnnotations struct {
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	PreferredFish []string  `json:"preferredFish"`
	CreatedAt     time.Time `json:"createdAt"`
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

	user2 := &UserWithAnnotations{
		Name:      "Tommy the Shark",
		Password:  "otorosushiisthebest",
		CreatedAt: time.Now(),
	}

	outputOfUser2, err := json.MarshalIndent(user2, "", " ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// cast the byte[] into a string
	fmt.Print(string(outputOfUser2))
}
