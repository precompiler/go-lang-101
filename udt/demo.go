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
}

type User struct {
	Name    string
	Email   string
	IsAdmin bool
}

type Meter float64
type Foot float64
