// This exercise will reinforce our understanding of method sets:
// ● create a type person struct
// ● attach a method speak to type person using a pointer receiver
// ○ *person
// ● create a type human interface
// ○ to implicitly implement the interface, a human must have the speak method
// ● create func “saySomething”
// ○ have it take in a human as a parameter
// ○ have it call the speak method
// ● show the following in your code
// ○ you CAN pass a value of type *person into saySomething
// ○ you CANNOT pass a value of type person into saySomething
// ● here is a hint if you need some help
// ○ https://play.golang.org/p/FAwcQbNtMG
package main

import "fmt"

type Person struct {
	name string
}

type Human interface {
	speak()
}

func (p *Person) speak() {
	fmt.Println(p.name + " is speaking!")
}

func saySomething(h Human) {
	h.speak()
}

func main() {
	mani := Person{"Subramanyam"}

	saySomething(&mani)

	//error
	// saySomething(mani)
}
