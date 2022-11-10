// Create TYPED and UNTYPED constants. Print the values of the constants

package main

import "fmt"

func main() {
	//untyped constants
	const (
		a = 1
		b = "Untyped Constant"
		c = true
	)

	//typed constants
	const (
		x int    = 10
		y string = "Typed Constant"
		z bool   = false
	)

	fmt.Println(a, b, c)
	fmt.Println(x, y, z)
}
