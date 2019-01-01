package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2)

	wg.Add(2)
	// fmt.Println("Start...")
	// go func() {
	// 	defer wg.Done()
	// 	for count := 0; count < 3; count++ {
	// 		for char := 'a'; char < 'a'+26; char++ {
	// 			fmt.Printf("%c ", char)
	// 		}
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	for count := 0; count < 3; count++ {
	// 		for char := 'A'; char < 'A'+26; char++ {
	// 			fmt.Printf("%c ", char)
	// 		}
	// 	}
	// }()
	// wg.Wait()
	// fmt.Println("\nCompleted...")
	go printPrime("A")
	go printPrime("B")
	wg.Wait()
	fmt.Println("Completed...")
}

func printPrime(id string) {
	defer wg.Done()
next:
	for i := 2; i < 5000; i++ {
		for k := 2; k < i; k++ {
			if i%k == 0 {
				continue next
			}
		}
		fmt.Printf("%s: %d\n", id, i)
	}
}
