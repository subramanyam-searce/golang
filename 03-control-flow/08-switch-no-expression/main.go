// Create a program that uses a switch statement with no switch expression specified.

package main

import "fmt"

func main() {
	switch {
	case 4 > 5:
		fmt.Println("4 is greater than 5")
	case 6 > 4:
		fmt.Println("6 is greater than 4")
	case 11 > 10:
		fmt.Println("5 is greater than 10")
	default:
		fmt.Println("Default")
	}
}
