package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter("A")
	go incCounter("B")
	wg.Wait()
	fmt.Println(counter)
}

func incCounter(id string) {
	defer wg.Done()
	for count := 0; count < 2; count++ {

		// value := counter
		// runtime.Gosched()
		// value++
		// counter = value

		atomic.AddInt64(&counter, 1)
	}
}
