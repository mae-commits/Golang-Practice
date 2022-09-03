package main

import (
	"fmt"
	"os"
)

func main() {
	// print os.Args' slice.
	// os.Args works as getting parameters from command line variables, like go run filename.go a b c"
	fmt.Println(os.Args)
	fmt.Println(os.Args[0]) // index 0 means the name of command.
	fmt.Println(os.Args[1]) // indices more than 0 means the arguments in a command line.
}
