// get this code working:
// ○ using func literal, aka, anonymous self-executing func
// ■ solution: https://play.golang.org/p/SHr3lpX4so
// ○ a buffered channel
// ■ solution: https://play.golang.org/p/Y0Hx6IZc3U

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	//anonymous function
	go func() {
		time.Sleep(time.Second * 3)
		c <- 42
	}()
	fmt.Println(<-c)

	//buffered channel
	bc := make(chan int, 1)
	bc <- 54
	fmt.Println(<-bc)
}
