// Start with this code. Create a custom error message using “fmt.Errorf”.

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, _ := json.Marshal(a)

	return bs, fmt.Errorf("Error Occurred: Custom Error")
}
