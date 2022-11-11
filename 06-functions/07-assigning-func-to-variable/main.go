// Assign a func to a variable, then call that func

package main

import "fmt"

func main() {
	myFunc := func() {
		fmt.Println("This function is assigned to a variable")
	}

	myFunc()
}
