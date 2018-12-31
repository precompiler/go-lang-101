package main

import (
	"fmt"
)

func main() {
	var user User
	fmt.Println(user)
	user2 := User{
		Name:    "Scott Tiger",
		Email:   "scott.tiger@oracle.com",
		IsAdmin: true, //trailing comma is a must
	}

	user3 := User{Name: "Scott Tiger", Email: "scott.tiger@oracle.com", IsAdmin: true}
	fmt.Println(user2)
	fmt.Println(user3)
	rect := Rectangle{}
	rect.SetWidth(10)
	rect.SetLength(20)
	fmt.Println(rect.Area())
	rect.dummy(0)
	fmt.Println(rect.Area())
}

type User struct {
	Name    string
	Email   string
	IsAdmin bool
}

type Meter float64
type Foot float64

type Rectangle struct {
	width  float32
	length float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (r *Rectangle) SetWidth(width float32) {
	r.width = width
}

func (r *Rectangle) SetLength(length float32) {
	r.length = length
}

func (r Rectangle) dummy(length float32) {
	r.length = length //mutating a copy
}
