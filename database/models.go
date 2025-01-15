package database

import "encoding/json"

// Address struct for storing address information
type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

// User struct for storing user information
type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}
