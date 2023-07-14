package main

import (
	"fmt"
	"os"

	"github.com/sethvargo/go-password/password"
)

func main() {
	var length int
	fmt.Printf("Enter password length (must be an int): ")

	_, err := fmt.Scanln(&length)
	if err != nil {
		fmt.Printf("Error reading input: %+v", err)
		os.Exit(1)
	}

	fmt.Println()

	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	res, err := password.Generate(length, 10, 10, false, false)
	if err != nil {
		fmt.Printf("Error generating password: %+v", err)
		os.Exit(1)
	}
	fmt.Printf("Password (len: %d): %s\n", len(res), res)
}
