package main

import (
	"greeting"
	"greeting/deutsch" // import deutsch package in greeting package.
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.Hallo()
	deutsch.GutenTag()
}
