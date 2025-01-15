# Custom Database Library

This is a simple custom database implementation written in Go. It provides basic operations such as storing, reading, and deleting JSON-based records in a file system. You can integrate it into your Go projects to easily manage user data or any other JSON-based records.

## Installation

To use this library in your project:

1. Run the following command to install the package:
```bash
   go get github.com/haile-paa/Pa-DataBase
```
   ```bash
  

  
   package main
    import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/haile-paa/Pa-DataBase" // Import the custom database
   )

    func main() {
	// Initialize your custom database
	db, err := customdb.New("./data") // Specify the directory to store the data
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Define the HTTP server and routes
	http.HandleFunc("/add-user", func(w http.ResponseWriter, r *http.Request) {
		// Only accept POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse the incoming JSON body into a User struct
		var user customdb.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		// Write the user data to the custom database
		err = db.Write("users", user.Name, user)
		if err != nil {
			http.Error(w, "Failed to save user data", http.StatusInternalServerError)
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User added successfully"))
	})

	// Start the HTTP server
	fmt.Println("Server is running at http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
   }
```
 Make a POST Request to Add Data:
You can make a POST request to the /add-user endpoint from your frontend or using a tool like Postman.

For example, the user data would look like this in JSON format:
```
 

json
{
  "Name": "John",
  "Age": "30",
  "Contact": "1234567890",
  "Company": "Acme Corp",
  "Address": {
    "City": "New York",
    "State": "NY",
    "Country": "USA",
    "Pincode": "10001"
  }
}


 ```
  Run the following command in your terminal to start the server:
```bash
  go run main.go
```
```bash
Here’s what the directory structure of the project looks like:
├── data/                # Directory to store data
│   └── users/           # Folder for storing users' data
├── main.go              # Backend code to interact with the custom database
└── go.mod               # Go module file
```
