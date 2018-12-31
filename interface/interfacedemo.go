package main

import (
	"fmt"
)

func main() {
	amdCard := AMD{"Radon"}
	nvCard := Nvdia{"GTX1800"}
	// compilation error
	// if you implement an interface using a pointer receiver
	// then only pointers of that type implement the interface
	//check(amdCard)
	installVideoCard(&amdCard)
	installVideoCard(nvCard)
}

type VideoCard interface {
	getModel() string
	init()
}

type AMD struct {
	model string
}

type Nvdia struct {
	model string
}

func (a *AMD) getModel() string {
	return a.model
}

func (a *AMD) init() {
	fmt.Println("Initializing ", a.model)
}

func (n Nvdia) getModel() string {
	return n.model
}

func (n Nvdia) init() {
	fmt.Println("Initializing ", n.model)
}

func installVideoCard(card VideoCard) {
	card.init()
}
