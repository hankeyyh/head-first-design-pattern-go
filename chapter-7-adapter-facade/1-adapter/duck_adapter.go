package main

import (
	"math/rand"
)
type DuckAdapter struct {
	duck Duck
}

func NewDuckAdapter(duck Duck) *DuckAdapter {
	return &DuckAdapter{duck: duck}
}

func (d *DuckAdapter) Gobble() {
	d.duck.Quack()
}

func (d *DuckAdapter) Fly() {
	// duck fly longer than turkey, 
	// so only fly the duck on average one in five times
	if rand.Int() % 5 == 0 {
		d.duck.Fly()
	}
}