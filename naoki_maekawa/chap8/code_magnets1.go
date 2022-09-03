package main

import "fmt"

/*
"student" is defined as the type of struct,
which contains "name" as string type argument
and "grade" as float64 argument.
*/

type student struct {
	name  string
	grade float64
}

/*
The func "printInfo" works as follows:
it prints the name of student
and the grade of student in the struct "s".
*/
func printInfo(s student) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

func main() {
	var s student
	s.name = "Alonzo Cole"
	s.grade = 92.3
	printInfo(s)
}
