package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Add this new function.
func reportPanic() {
	// Call "recover" and store its return value.
	p := recover()
	// If "recover" returned nil, there is no panic...
	if p == nil {
		return
	}
	// Without "()", get the underlying "error" value...
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		// If the panic value is not an error, resume panicking with the same value.
		panic(p)
	}
}

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	// Before calling code that might panic, defer a call to our new reportPanic function.
	defer reportPanic()
	panic("some other issue")
	scanDirectory("my_directory")
}
