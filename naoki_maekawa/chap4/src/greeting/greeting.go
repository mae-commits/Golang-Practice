// it is one of packages for other main codes.
package greeting // to defibe as packages, declare individual name as package call line.

import "fmt"

func Hello() { // In order to call functions in file "greeting.go", define fucntions' name starting from uppercase.
	fmt.Println("Hello!")
}

func Hi() {
	fmt.Println("Hi!")
}
