package main

import (
	"sync"
)

type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

func (c *ChocolateBoiler) Fill() {
	if c.empty {
		c.empty = false
		c.boiled = false
		// fill the boiler with a milk/chocolate mixture
	}
}

func (c *ChocolateBoiler) Drain() {
	if !c.empty && c.boiled {
		// drain the boiled milk and chocolate
		c.empty = true
	}
}

func (c *ChocolateBoiler) Boil() {
	if !c.empty && !c.boiled {
		// bring the contents to a boil
		c.boiled = true
	}
}

var boiler *ChocolateBoiler
var mu sync.Mutex

func NewChocolateBoiler() *ChocolateBoiler {
	if boiler != nil {
		return boiler
	}
	mu.Lock()
	if boiler == nil {
		boiler = &ChocolateBoiler{
			empty:  true,
			boiled: false,
		}
	}
	mu.Unlock()
	return boiler
}
