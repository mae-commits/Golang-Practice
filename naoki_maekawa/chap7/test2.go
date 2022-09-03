package main

import (
	"fmt"
)

func main() {
	// Create a map and declare a variable to hold it.
	// make(map[key type] value type)
	ranks := make(map[string]int)
	/*
		"gold", "silver", "blonze" are keys.
	*/
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])
	/*
		Define keys and values as follows;
		map[key type] value type{
			key1 : value1,
			key2 : value2,
			...
		}
	*/
	elements := map[string]string{
		"H":  "Hydrogen",
		"Li": "Litium",
	}
	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])
}
