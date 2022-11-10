// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main

import "fmt"

func main() {

	const (
		y1 = iota + 2023
		y2
		y3
		y4
	)

	const (
		y5 = iota + 2023
		y6
	)

	fmt.Println(y1, y2, y3, y4, y5, y6)
}
