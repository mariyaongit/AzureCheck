package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Function to generate a unique identifier based on instance-specific information
func generateUniqueID() int {
	return int(time.Now().UnixNano() % 3)
}

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Could not get hostname", http.StatusInternalServerError)
		return
	}

	// Read the unique identifier
	uniqueID := generateUniqueID()

	// Customize text and color based on unique ID
	var text, fontColor string
	switch uniqueID {
	case 0:
		text = "Welcome to our site! Explore our range of services and find out how we can assist you"
		fontColor = "red"
	case 1:
		text = "Please review the details provided and complete the necessary actions to proceed."
		fontColor = "blue"
	case 2:
		text = "Thank you for your submission. We appreciate your input and will get back to you as soon as possible"
		fontColor = "green"
	}

	// Output HTML with font color, text, and hostname
	fmt.Fprintf(w, `<html>
    <head>
        <title>Hostname</title>
    </head>
    <body>
        <h1 style="color: %s;">%s</h1>
        <p>Hostname: %s</p>
        <p>Unique ID: %d</p>
    </body>
    </html>`, fontColor, text, hostname, uniqueID)
    }

func main() {
	http.HandleFunc("/", hostnameHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Could not start server: %s\n", err)
	}
}
