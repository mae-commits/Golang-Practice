package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

/*
If a type declares methods with pointer receivers,
then you will only be able to use pointers to that type
when assigning to interface variables.
*/
func main() {
	s := Switch("off")
	var t Toggleable = &s
	t.toggle()
	t.toggle()
}
