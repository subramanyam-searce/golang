// Write a program that
// ● assigns an int to a variable
// ● prints that int in decimal, binary, and hex
// ● shifts the bits of that int over 1 position to the left, and assigns that to a variable
// ● prints that variable in decimal, binary, and hex

package main

import "fmt"

func main() {
	a := 8 //1000 in binary

	fmt.Printf("%d, %b, %x\n", a, a, a)

	//left shift
	b := a << 1 //10000 in binary | 16 in Decimal
	fmt.Printf("%d, %b, %x, %T\n", b, b, b, b)

	//>> right shift => removes values at the end.
}
