//Create a variable of type string using a raw string literal. Print it.

package main

import "fmt"

func main() {
	str := `Hello
There
	this is a
	"quoted text"`

	fmt.Println(str)
}
