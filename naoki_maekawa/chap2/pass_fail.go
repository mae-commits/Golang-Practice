// pass_fail reports a grade is passing or falling.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader1 := bufio.NewReader(os.Stdin)
	// reader2 := bufio.NewReader(os.Stdin)
	input1, err := reader1.ReadString('\n')
	// input2, _ := reader2.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input1)

	// fmt.Print(input2)

}

/*

bufio means buffered.

bufio.Newreader returns a new bufio.Reader

os.Stdin is that the Reader will read from standard input (= Stdin) from the keyboard.

reader.ReadString method requires an argument with a rune(character) that marks the end of the input
when it returns what the user typed, as a string.

fmt.Print allow us to print message in the same line without skipping line unlike fmt.Println.

note: Golang returns any number of values, so when you input some from Stdin, you have to prepare more than two variables.
	  But if you'd like to use only the first value, define the second one as "_", which is called "Blank identifier".

*/
