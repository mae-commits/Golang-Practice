package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func d(writer http.ResponseWriter, request *http.Request) {
	write(writer, "z")
}

func e(writer http.ResponseWriter, request *http.Request) {
	write(writer, "x")
}

func f(writer http.ResponseWriter, request *http.Request) {
	write(writer, "y")
}
func main() {
	http.HandleFunc("/a", f)
	http.HandleFunc("/b", d)
	http.HandleFunc("/c", e)
	err := http.ListenAndServe("localhost:4567", nil)
	log.Fatal(err)
}
