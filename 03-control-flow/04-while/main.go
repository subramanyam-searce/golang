// Create a for loop using this syntax
// ● for { }
// Have it print out the years you have been alive

package main

import "fmt"

func main() {
	i := 2001
	currentYear := 2022

	for {
		if i > currentYear {
			break
		}
		fmt.Println(i)
		i += 1
	}

}
