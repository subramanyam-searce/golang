// Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
// has a value of type error as a parameter. Create a value of type “customErr” and pass it into
// “foo”. If you need a hint, here is one.

package main

import "fmt"

type customErr struct {
	s string
}

//This is already built-in
// type error interface {
// 	Error() string
// }

func foo(e error) {
	fmt.Printf("%T, %v", e, e)
}

func (e customErr) Error() string {
	return "Custom Error Function"
}

func main() {
	ce := customErr{"custom Err"}
	foo(ce)
}
