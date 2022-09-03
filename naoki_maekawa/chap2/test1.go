package main

import ( // import packages
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now() // now is a function of time.Time
	var year int = now.Year()      // year is the method
	fmt.Println(year)
}

/*

"time" is the package that has a Time type that represents a date(year,month,and day)
and time(hour,minute,second,etc...)

time.Time value has Year method that returns the year.
*/
