package main

import "fmt"

func negate(myBoolean *bool) { // *bool: define as boolean pointer
	*myBoolean = !*myBoolean // *myBoolean: define myBoolean as pointer type of variables
}

func main() {
	truth := true
	negate(&truth) // input truth's addless 
	fmt.Println(truth)
	lies := false
	negate(&lies) // input lies' addless 
	fmt.Println(lies)
}
