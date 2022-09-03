package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Call http.Get with URL we want to retrieve.
	response, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	// Release the network connection once the "main" function exits.
	defer response.Body.Close()
	// Read all the data in the response.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Convert the data to a string and print it.
	fmt.Println(string(body))
}
