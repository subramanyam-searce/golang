// write a program that
// ○ launches 10 goroutines
// ■ each goroutine adds 10 numbers to a channel
// ○ pull the numbers off the channel and print them

//solutions
// ● https://play.golang.org/p/R-zqsKS03P
// ● https://play.golang.org/p/quWnlwzs2z
// ● https://go.dev/play/p/WqYnBC_CiKn

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 10; j < 20; j++ {
				c <- j
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
}
