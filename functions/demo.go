package main

import "fmt"

func main() {
	a := 4
	b := 2
	println(doMath(add, a, b))
	println(doMath(minus, a, b))
	println(doMath(mul, a, b))
	println(doMath(log(div), a, b))
}

func doMath(logic func(int, int) int, a int, b int) int {
	return logic(a, b)
}

func add(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func log(f func(int, int) int) func(int, int) int {
	return func(a int, b int) int {
		fmt.Println("a: ", a, " b: ", b)
		return f(a, b)
	}
}
