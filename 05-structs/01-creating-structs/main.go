// Create your own type “person” which will have an underlying type of “struct” so that it can store
// the following data:
// ● first name
// ● last name
// ● favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
// which stores the favorite flavors.

package main

import "fmt"

type Person struct {
	firstName          string
	lastName           string
	favIcecreamFlavors []string
}

func main() {
	subbu := Person{
		firstName:          "Subramanyam",
		lastName:           "R",
		favIcecreamFlavors: []string{"Strawberry", "Choculate", "Butterscotch"},
	}

	akash := Person{
		firstName:          "Akash",
		lastName:           "Anilkumar",
		favIcecreamFlavors: []string{"Pista", "Vanilla", "Mango"},
	}

	fmt.Println("Person: ", subbu.firstName, subbu.lastName)
	for _, v := range subbu.favIcecreamFlavors {
		fmt.Println(v)
	}

	fmt.Println("Person: ", akash.firstName, akash.lastName)
	for _, v := range akash.favIcecreamFlavors {
		fmt.Println(v)
	}

}
