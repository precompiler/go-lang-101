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

	order := []string{
		"Coffee",
		"Cream",
		"Cream",
	}
	fmt.Println(CalcPrice(order))

	order = []string{
		"Coffee",
		"Coffee",
		"Cocoa",
		"Cream",
	}
	fmt.Println(CalcPrice(order))
}

type Employee struct {
	FirstName string
	LastName  string
	EmpID     int
}

type Condiment interface {
	GetPrice() float32
	New(Condiment) Condiment
}

type Cream struct {
	Price float32
	Cond  Condiment
}

type Sugar struct {
	Price float32
	Cond  Condiment
}

type Cocoa struct {
	Price float32
	Cond  Condiment
}

func (c *Cream) GetPrice() float32 {
	if c.Cond != nil {
		return c.Cond.GetPrice() + c.Price
	} else {
		return c.Price
	}
}

func (s *Sugar) GetPrice() float32 {
	if s.Cond != nil {
		return s.Cond.GetPrice() + s.Price
	} else {
		return s.Price
	}
}

func (c *Cocoa) GetPrice() float32 {
	if c.Cond != nil {
		return c.Cond.GetPrice() + c.Price
	} else {
		return c.Price
	}
}

func (c *Cream) New(cond Condiment) Condiment {
	c.Price = 0.5
	c.Cond = cond
	return c
}

func (s *Sugar) New(cond Condiment) Condiment {
	s.Price = 0.2
	s.Cond = cond
	return s
}

func (c *Cocoa) New(cond Condiment) Condiment {
	c.Price = 0.8
	c.Cond = cond
	return c
}

type Coffee struct {
	Price float32
	Cond  Condiment
}

func (c *Coffee) GetPrice() float32 {
	if c.Cond != nil {
		return c.Cond.GetPrice() + c.Price
	} else {
		return c.Price
	}
}

func (c *Coffee) New(cond Condiment) Condiment {
	c.Price = 2
	c.Cond = cond
	return c
}

type Zero struct {
	Price float32
	Cond  Condiment
}

func (z *Zero) GetPrice() float32 {
	if z.Cond != nil {
		return z.Cond.GetPrice() + z.Price
	} else {
		return z.Price
	}
}

func (z *Zero) New(cond Condiment) Condiment {
	z.Price = 0
	z.Cond = cond
	return z
}

func CalcPrice(condiments []string) float32 {
	condimentRegistry := map[string]reflect.Type{
		"Coffee": reflect.TypeOf(Coffee{}),
		"Cocoa":  reflect.TypeOf(Cocoa{}),
		"Sugar":  reflect.TypeOf(Sugar{}),
		"Cream":  reflect.TypeOf(Cream{}),
	}
	var preCond Condiment = nil
	for _, cond := range condiments {
		if condType, exists := condimentRegistry[cond]; exists {
			c := reflect.New(condType)
			if preCond == nil {
				c.MethodByName("New").Call([]reflect.Value{reflect.ValueOf(&Zero{})})
			} else {
				c.MethodByName("New").Call([]reflect.Value{reflect.ValueOf(preCond)})
			}
			preCond = c.Interface().(Condiment)
			// v := reflect.ValueOf(preCond)
			// fmt.Println(v)
			// if !v.IsNil() {
			// 	c.Elem().FieldByName("Cond").Set(v)
			// }
			// preCond = c.Elem().Interface().(Condiment)
		}
	}
	return preCond.GetPrice()
}
