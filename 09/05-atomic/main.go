// Fix the race condition you created in exercise #4 by using package atomic

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64
	wg := sync.WaitGroup{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&count))
}
