// Create and use an anonymous struct.

package main

import "fmt"

func main() {
	myCar := struct {
		model  string
		wheels int
	}{
		model:  "Aston Martin",
		wheels: 4,
	}

	fmt.Println(myCar)
}
