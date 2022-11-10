// Building on the previous hands-on exercise, create a program that uses “else if” and “else”

package main

import "fmt"

func main() {
	if 4 > 5 {
		fmt.Println("4 is greater than 5")
	} else if 4 > 10 {
		fmt.Println("4 is greater than 10")
	} else {
		fmt.Println("4 is less than both 5 and 10")
	}
}
