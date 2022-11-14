// ● Using goroutines, create an incrementer program
// ○ have a variable to hold the incrementer value
// ○ launch a bunch of goroutines
// ■ each goroutine should
// ● read the incrementer value
// ○ store it in a new variable
// ● yield the processor with runtime.Gosched()
// ● increment the new variable
// ● write the value in the new variable back to the incrementer
// variable
// ● use waitgroups to wait for all of your goroutines to finish
// ● the above will create a race condition.
// ● Prove that it is a race condition by using the -race flag
// ● if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0
	wg := sync.WaitGroup{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			val := count
			runtime.Gosched()
			val += 1
			count = val
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
