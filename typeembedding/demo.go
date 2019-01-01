package main

import "fmt"

func main() {
	benz := Benz{
		Vehicle{
			brand: "Benz",
		},
	}

	benz.start()
	benz.Vehicle.start()
}

type Vehicle struct {
	numOfWheel int
	brand      string
}

func (v *Vehicle) start() {
	fmt.Println("Starting the vehicle")
}

type Benz struct {
	Vehicle
}

func (b *Benz) start() {
	fmt.Println("Starting ", b.brand)
}
