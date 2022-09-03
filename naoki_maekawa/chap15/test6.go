package main

import "fmt"

// We can tell from the calFunction body that the passed function accepts no parameters.
func callFunction(passedFunction func()) {
	passedFunction()
}

// We can tell from the callTwice body that the passed function accepts no parameters.
func callTwice(passedFunction func()) {
	passedFunction()
	passedFunction()
}

// We can tell from the callWithArguments body
// that the passed-in function must accept these parameter types.
func callWithArguments(passedFunction func(string, bool)) {
	passedFunction("This sentence is", false)
}

func printReturnValue(passedFunction func() string) {
	fmt.Println(passedFunction())
}

func functionA() {
	fmt.Println("function called")
}

func functionB() string {
	fmt.Println("function called")
	return "Returning from function"
}

func functionC(a string, b bool) {
	fmt.Println("function called")
	fmt.Println(a, b)
}

func main() {
	callFunction(functionA)
	callTwice(functionA)
	callWithArguments(functionC)
	printReturnValue(functionB)

}
