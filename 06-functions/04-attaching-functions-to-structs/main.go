// ● Create a user defined struct with
// ○ the identifier “person”
// ○ the fields:
// ■ first
// ■ last
// ■ age
// ● attach a method to type person with
// ○ the identifier “speak”
// ○ the method should have the person say their name and age
// ● create a value of type person
// ● call the method from the value of type person

package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) speak() {
	fmt.Println("Hello, My name is", p.first, p.last+". Iam", p.age, "years old.")
}

func main() {
	subbu := Person{
		first: "Subramanyam",
		last:  "R",
		age:   21,
	}

	subbu.speak()
}
