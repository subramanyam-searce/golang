// ○ create a func with the identifier foo that returns an int
// ○ create a func with the identifier bar that returns an int and a string
// ○ call both funcs
// ○ print out their results

package main

import "fmt"

func foo() int {
	return 0
}

func bar() (int, string) {
	return 3, "Mani"
}

func main() {
	foo_int := foo()

	bar_int, bar_string := bar()

	fmt.Println(foo_int, bar_int, bar_string)
}
