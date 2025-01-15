# Custom Database Library

This is a simple custom database implementation written in Go. It provides basic operations such as storing, reading, and deleting JSON-based records in a file system.

## Installation

To use this library in your project:

1. Run `go get github.com/haile-paa/Pa-DataBase`
2. Import the `custom-db` package in your project.

## Usage

```go
package main

import (
	"fmt"
	"log"
	"github.com/haile-paa/Pa-DataBase"
)

func main() {
	// Initialize your custom database
	db, err := customdb.New("./data")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Example of writing a user record
	user := customdb.User{
		Name:    "John",
		Age:     "30",
		Contact: "1234567890",
		Company: "Acme Corp",
		Address: customdb.Address{"New York", "NY", "USA", "10001"},
	}
	if err := db.Write("users", user.Name, user); err != nil {
		fmt.Println("Error writing user:", err)
	}

	// Example of reading a user record
	var readUser customdb.User
	if err := db.Read("users", "John", &readUser); err != nil {
		fmt.Println("Error reading user:", err)
	} else {
		fmt.Printf("Read user: %+v\n", readUser)
	}

	// Example of deleting a user record
	if err := db.Delete("users", "John"); err != nil {
		fmt.Println("Error deleting user:", err)
	} else {
		fmt.Println("User deleted successfully")
	}
}
