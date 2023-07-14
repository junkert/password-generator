package main

import (
	"fmt"

	"github.com/sethvargo/go-password/password"
)

func main() {
	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	res, err := password.Generate(32, 10, 10, false, false)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Printf("Password: ", res)
}
