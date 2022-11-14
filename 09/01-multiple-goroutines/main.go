// ● in addition to the main goroutine, launch two additional goroutines
// ○ each additional goroutine should print something out
// ● use waitgroups to make sure each goroutine finishes before your program exists

package main

import (
	"fmt"
	"sync"
)

func foo() {
	fmt.Println("Function Foo")
	wg.Done()
}

func bar() {
	fmt.Println("Function Bar")
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
