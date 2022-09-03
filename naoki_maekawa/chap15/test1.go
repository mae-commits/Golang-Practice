package main

import (
	"log"
	// Import the "net/http" package.
	"net/http"
)

// A value for updating the response that will be sent to the browser.
// A value representing the request from the browser.
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	// Add "Hello, web!" to the response.
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// If we receive a request for a URL ending in "/hello"...
	// then call the viewHandler function to generate a response.
	http.HandleFunc("/hello", viewHandler)
	// Listen for browser requests, and respond to them.
	// Host name is "localhost", Port number is "8080"
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
