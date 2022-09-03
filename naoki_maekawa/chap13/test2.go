package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// Get the sizes of several pages.
	// Convert the responseSize calls to Go statements.
	go responseSize("https://example.com/")
	go responseSize("https://golang.org/")
	go responseSize("https://golang.org/doc")
	// Sleep for 5 seconds.
	time.Sleep(5 * time.Second)
}

// Take the URL as a parameter.
func responseSize(url string) {
	// Move the code that gets the page to a separate function.
	// Print which URL we are retrieving.
	fmt.Println("Getting", url)
	// Get the given URL.
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	/*
		The size of the slice of bytes is
		the same as the size of the page.
	*/
	fmt.Println(len(body))
}
