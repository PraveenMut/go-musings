package main

import (
	"time"
)

// User struct takes name, password and preferredfish fields to create a user profile
type User struct {
	Name          string
	password      string
	preferredFish []string
	createdAt     time.Time
}

func main() {

}
