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

// Channel we pass to responseSize will carry Pages, not ints.
func responseSize(url string, channel chan Page) {
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
	// Send back a Page with both the current URL and the page size.
	channel <- Page{URL: url, Size: len(body)}
}

func main() {
	// Make a channel of int values.
	pages := make(chan Page)
	urls := []string{"https://example.com", "https://golang.org/doc", "https://golang.org"}
	for _, url := range urls {
		// Pass the channel to responseSize.
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		// Receive the Page.
		page := <-pages
		// Print its URL and size together.
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}
