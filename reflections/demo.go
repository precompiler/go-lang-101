package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 1
	ref := reflect.ValueOf(num)
	fmt.Println(ref.Type(), ref)

	emp := Employee{
		FirstName: "Scott",
		LastName:  "Tiger",
		EmpID:     1,
	}
	fmt.Println(emp)

	empRef := reflect.ValueOf(&emp)
	fmt.Println(empRef.Type(), empRef.Elem().NumMethod(), empRef.Elem().NumField())
	for i := 0; i < empRef.Elem().NumField(); i++ {
		fmt.Print(empRef.Elem().Type().Field(i).Name, " ", empRef.Elem().Field(i).Type().Name(), " ", empRef.Elem().Field(i).Interface(), ";")
	}
	fmt.Println()
	empRef.Elem().FieldByName("FirstName").SetString("Dummy")
	fmt.Println(emp)
}

type Employee struct {
	FirstName string
	LastName  string
	EmpID     int
}
