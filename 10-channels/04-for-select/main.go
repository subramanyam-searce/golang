// Starting with this code, pull the values off the channel using a select statement

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("From Channel C:", v)
		case v := <-q:
			fmt.Println("From Channel Q:", v)
			return
		}
	}
}

func gen(q chan int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 0
	}()
	return c
}
