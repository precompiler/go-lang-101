package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	pipe := make(chan string, 10)
	cook := Cook{&pipe}
	busboy := Busboy{&pipe}
	go cook.makeDish()
	go busboy.deliverDish()
	wg.Wait()
}

type Cook struct {
	pipe *chan string
}

type Busboy struct {
	pipe *chan string
}

func (c *Cook) makeDish() {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		dish := fmt.Sprintf("%s%d", "Dish", i)
		fmt.Println("Dish cooked: ", dish)
		*(c.pipe) <- dish
		time.Sleep(10 * time.Millisecond)
	}
}

func (b *Busboy) deliverDish() {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		dish := <-*(b.pipe)
		fmt.Println("Delivering dish: ", dish)
		time.Sleep(1000 * time.Millisecond)
	}
}
