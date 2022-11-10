// 1. At the package level scope, assign the following values to the three variables
// a. for x assign 42
// b. for y assign “James Bond”
// c. for z assign true
// 2. in func main
// Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 26
// a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
// returned value of TYPE string using the short declaration operator to a
// VARIABLE with the IDENTIFIER “s”
// b. print out the value stored by variable “s”

package main

import "fmt"

// 1a b c
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//2a
	s := fmt.Sprintf("%v %v %v", x, y, z)

	//2b
	fmt.Println(s)
}
