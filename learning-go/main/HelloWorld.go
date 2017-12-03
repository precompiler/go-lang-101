package main

import (
	"fmt"
	"math"
)
//or
//import "fmt"
//import "math"
func add(a int, b int) int {
	return a + b
}
func multiReturn() (string, int, float64) {
	return "Hello", 1, 1.1
}
func main() {
	fmt.Println("Hello World")
	fmt.Println(math.Pi)
	fmt.Println(add(1, 1))
	a, b, c := multiReturn()
	fmt.Println(a, b, c)
	fmt.Println(multiReturn())
	sum := add(1, 9)
	fmt.Println(sum)

	var x int = 10
	fmt.Println(x)
	var s = add(100, 100)
	fmt.Println(s)

	var aa, bb = 1, true
	fmt.Println(aa, bb)

	fmt.Println(float64(256))

	const hello = "Hello"
	fmt.Println(hello)
}
