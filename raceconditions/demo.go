package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter  int64
	counter2 int64
	wg       sync.WaitGroup
	wg2      sync.WaitGroup
	mutex    sync.Mutex
)

func main() {
	wg.Add(2)
	go incCounter("A")
	go incCounter("B")
	wg.Wait()
	fmt.Println(counter)

	wg2.Add(2)
	go incCounter2("C")
	go incCounter2("D")
	wg2.Wait()
	fmt.Println(counter2)
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

func incCounter2(id string) {
	defer wg2.Done()
	for count := 0; count < 2; count++ {
		mutex.Lock()
		value := counter2
		runtime.Gosched()
		value++
		counter2 = value
		mutex.Unlock()
	}
}
