// Create a program that prints out your OS and ARCH. Use the following commands to run it
// ● go run
// ● go build
// ● go install

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOOS: ", runtime.GOOS)
	fmt.Println("GOARCH: ", runtime.GOARCH)
}
