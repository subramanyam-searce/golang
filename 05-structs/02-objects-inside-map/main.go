// Take the code from the previous exercise, then store the values of type person in a map with the
// key of last name. Access each value in the map. Print out the values, ranging over the slice.

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

	mp := map[string]Person{
		subbu.lastName: subbu,
		akash.lastName: akash,
	}

	for _, v := range mp {
		fmt.Println("Person: "+v.firstName, v.lastName)
		for _, fif := range v.favIcecreamFlavors {
			fmt.Println(fif)
		}
		fmt.Println("-------------")
	}

}
