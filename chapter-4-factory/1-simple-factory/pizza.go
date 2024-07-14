package main

import (
	"fmt"
)

type Pizza interface {
	prepare()
	bake()
	cut()
	box()
	toString()
	getName() string
}


type BasicPizza struct {
	name string
	dough string
	sauce string
	toppings []string
}

func (b *BasicPizza) prepare() {
	fmt.Println("Prepare " + b.name)
}

func (b *BasicPizza) bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (b *BasicPizza) cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (b *BasicPizza) box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (b *BasicPizza) toString() {
	fmt.Println("------- ", b.name, " -------")
	fmt.Println(b.dough)
	fmt.Println(b.sauce)
	for _, topping := range b.toppings {
		fmt.Println(topping)
	}
}

func (b *BasicPizza) getName() string {
	return b.name
}