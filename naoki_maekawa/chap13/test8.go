package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

// We will pass a channel in to responseSize for it to send page sizes over.
func responseSize(url string, channel chan int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Instead of returning the page size, send it over the channel.
	channel <- len(body)
}

func main() {
	// Make a channel of int values.
	sizes := make(chan int)
	// Pass the channel with each call to responseSize.
	go responseSize("https://example.com/", sizes)
	go responseSize("https://golang.org/", sizes)
	go responseSize("https://golang.org/doc", sizes)
	// There will be three sends on the channel, so do three receives.
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
}
