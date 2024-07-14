package main

import (
	"fmt"

	"github.com/hankeyyh/head-first-design-pattern-go/chapter-4-factory/3-abstract-factory/ingredient"
)

type Pizza interface {
	prepare()
	bake()
	cut()
	box()
	toString()
	getName() string
	setName(name string)
}

type BasicPizza struct {
	name      string
	dough     ingredient.Dough
	sauce     ingredient.Sauce
	veggies   []ingredient.Veggies
	cheese    ingredient.Cheese
	pepperoni ingredient.Pepperoni
	clams     ingredient.Clams
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
	if b.dough != nil {
		fmt.Println(b.dough.ToDoughString())
	}
	if b.sauce != nil {
		fmt.Println(b.sauce.ToSauceString())
	}
	if b.cheese != nil {
		fmt.Println(b.cheese.ToCheeseString())
	}
	for _, veggie := range b.veggies {
		fmt.Println(veggie.ToVeggiesString())
	}
	if b.clams != nil {
		fmt.Println(b.clams.ToClamsString())
	}
	if b.pepperoni != nil {
		fmt.Println(b.pepperoni.ToPepperoniString())
	}
}

func (b *BasicPizza) getName() string {
	return b.name
}

func (b *BasicPizza) setName(name string) {
	b.name = name
}

type CheesePizza struct {
	*BasicPizza
}

func NewCheesePizza(factory ingredient.IngredientFactory) *CheesePizza {
	return &CheesePizza{
		BasicPizza: &BasicPizza{
			dough:  factory.CreateDough(),
			sauce: factory.CreateSauce(),
			cheese: factory.CreateCheese(),
		},
	}
}

type ClamPizza struct {
	*BasicPizza
}

func NewClamPizza(factory ingredient.IngredientFactory) *ClamPizza {
	return &ClamPizza{
		BasicPizza: &BasicPizza{
			dough:  factory.CreateDough(),
			sauce: factory.CreateSauce(),
			cheese: factory.CreateCheese(),
			clams: factory.CreateClams(),
		},
	}
}

type PepperoniPizza struct {
	*BasicPizza
}

func NewPepperoniPizza(factory ingredient.IngredientFactory) *PepperoniPizza {
	return &PepperoniPizza{
		BasicPizza: &BasicPizza{
			dough:  factory.CreateDough(),
			sauce: factory.CreateSauce(),
			cheese: factory.CreateCheese(),
			veggies: factory.CreateVeggies(),
			pepperoni: factory.CreatePepperoni(),
		},
	}
}

type VeggiePizza struct {
	*BasicPizza
}

func NewVeggiePizza(factory ingredient.IngredientFactory) *VeggiePizza {
	return &VeggiePizza{
		BasicPizza: &BasicPizza{
			dough:  factory.CreateDough(),
			sauce: factory.CreateSauce(),
			cheese: factory.CreateCheese(),
			veggies: factory.CreateVeggies(),
		},
	}
}