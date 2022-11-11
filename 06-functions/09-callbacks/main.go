// A “callback” is when we pass a func into a func as an argument. For this exercise,
// ● pass a func into a func as an argument

package main

import "fmt"

func foo(callback func(a int) int) {
	fmt.Println(callback(5))
}

func main() {
	foo(func(a int) int {
		return a * a
	})
}
