// Create a program that uses a switch statement with the switch expression specified as a
// variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	favSport := "Cricket"

	switch favSport {
	case "Football":
		fmt.Println("Favourite Sport is Football")
	case "Basketball":
		fmt.Println("Favourite Sport is Basketball")
	case "Cricket":
		fmt.Println("Favourite Sport is Cricket")
	case "Hockey":
		fmt.Println("Favourite Sport is Hockey")
	default:
		fmt.Println("No Favourite Sport")
	}

}
