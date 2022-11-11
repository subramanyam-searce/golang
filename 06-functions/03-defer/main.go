// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

package main

import "fmt"

func foo() {
	defer fmt.Println("Deferred Statement")
	//some Operation
	fmt.Println("Foo Funtion running")
	fmt.Println("Returning Function")
	return
}

func main() {
	foo()
}
