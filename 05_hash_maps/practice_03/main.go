package main

import (
	"fmt"
)

func main() {
	type User struct {
		Age int
	}

	users := make(map[string]User)

	users["Alice"] = User{Age: 25}

	fmt.Println(users)

	// users["Alice"].Age = 26

	// ptr := &users["Alice"]

	userAlice := users["Alice"]

	userAlice.Age = 26

	users["Alice"] = userAlice

	fmt.Println(users)

}
