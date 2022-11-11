// ● Create a func which returns a func
// ● assign the returned func to a variable
// ● call the returned func

package main

import "fmt"

func foo() func() {
	return func() {
		fmt.Println("This is a function, returned from another function (foo)")
	}
}

func main() {
	myVar := foo()

	myVar()
}
