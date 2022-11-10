// Building on the code from the previous example
// 1. at the package level scope, using the “var” keyword, create a VARIABLE with the
// IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom
// TYPE “x”
// a. eg:
// 2. in func main
// a. this should already be done
// i. print out the value of the variable “x”
// ii. print out the type of the variable “x”
// iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
// Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 27
// iv. print out the value of the variable “x”
// b. now do this
// i. now use CONVERSION to convert the TYPE of the VALUE stored in “x”
// to the UNDERLYING TYPE
// 1. then use the “=” operator to ASSIGN that value to “y”
// 2. print out the value stored in “y”
// 3. print out the type of “y”

package main

import "fmt"

type MyInt int

var x MyInt

// 1a
var y int

func main() {
	//2a
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	//2b1
	y = int(x)

	//2b2
	fmt.Println(y)

	//2b3
	fmt.Printf("%T\n", x)
}
