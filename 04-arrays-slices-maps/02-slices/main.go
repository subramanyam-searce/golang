// ● Using a COMPOSITE LITERAL:
// ● create a SLICE of TYPE int
// ● assign 10 VALUES
// ● Range over the slice and print the values out.
// ● Using format printing
// ○ print out the TYPE of the slice

package main

import "fmt"

func main() {
	//composite literal
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range slc {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", slc)
}
