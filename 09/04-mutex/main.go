// Fix the race condition you created in the previous exercise by using a mutex
// ● it makes sense to remove runtime.Gosched()

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			val := count
			val += 1
			count = val
			wg.Done()
			mu.Unlock()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(count)
}
