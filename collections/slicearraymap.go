package main

import "fmt"

func main() {
	var integers [5]int
	fmt.Println(integers)
	integers[1] = 1
	fmt.Println(integers)
	strings := [3]string{"Hello", "World", "!"}
	for _, item := range strings {
		fmt.Print(item)
	}
	fmt.Println()
	fmt.Println([...]int{1, 2, 3})
	fmt.Println([]int{0: 10, 8: 20})
	fmt.Println([10]int{0: 10, 8: 20})

	words := [5]*string{4: new(string)}
	fmt.Println(words)
	a, b := "a", "b"
	words[0] = &a
	words[1] = &b
	fmt.Println(words)
	fmt.Println(*words[0])

	colors1 := []string{"red", "green", "blue"}
	colors2 := colors1
	colors4 := [3]string{}
	colors3 := [3]string{"red", "green", "blue"}
	colors4 = colors3
	colors2[0] = "black"
	fmt.Println(colors1[0], colors2[0])
	colors4[0] = "white"
	fmt.Println(colors3[0], colors4[0])

	slice := make([]string, 5)
	fmt.Println(slice)

	num1 := []int{1, 2, 3}
	var num2 []int
	num2 = num1
	addr1 := &num1[0]
	addr2 := &num2[0]
	fmt.Println(&num2, &num1, addr1, addr2)

	var stringArray []string = make([]string, 2)
	stringArray[0] = "a"
	stringArray[1] = "b"
	printStringArray(&stringArray)

	var stringPointerArray []*string = make([]*string, 2) //array of string pointers
	msg := "Hello"
	stringPointerArray[0] = &msg
	fmt.Println(*stringPointerArray[0])

}

func printStringArray(strings *[]string) { // a pointer pointed to a string slice
	fmt.Println("=======")
	for _, item := range *strings {
		fmt.Println(item)
	}
	fmt.Println("=======")
}
