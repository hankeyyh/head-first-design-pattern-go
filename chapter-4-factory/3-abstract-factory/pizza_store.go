package main

import "github.com/hankeyyh/head-first-design-pattern-go/chapter-4-factory/3-abstract-factory/ingredient"

type PizzaStore interface {
	createPizza(pizzaType string) Pizza
	orderPizza(pizzaType string) Pizza
}

type BasicPizzaStore struct {
	createPizza func(pizzaType string) Pizza
}

func (d *BasicPizzaStore) orderPizza(pizzaType string) Pizza {
	pizza := d.createPizza(pizzaType)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
	return pizza
}

type NewYorkPizzaStore struct {
	BasicPizzaStore
}

func NewNewYorkPizzaStore() *NewYorkPizzaStore {
	result := new(NewYorkPizzaStore)
	result.BasicPizzaStore.createPizza = result.createPizza
	return result
}

func (n *NewYorkPizzaStore) createPizza(pizzaType string) Pizza {
	ingredientFactory := ingredient.NewNYIngredientFactory()
	var pizza Pizza
	switch pizzaType {
	case "cheese":
		pizza = NewCheesePizza(ingredientFactory)
		pizza.setName("NY Style Cheese Pizza")
	case "veggie":
		pizza =  NewVeggiePizza(ingredientFactory)
		pizza.setName("NY Style Veggie Pizza")
	case "clam":
		pizza =  NewClamPizza(ingredientFactory)
		pizza.setName("NY Style Clam Pizza")
	case "pepperoni":
		pizza =  NewPepperoniPizza(ingredientFactory)
		pizza.setName("NY Style Pepperoni Pizza")
	}
	return pizza
}

type ChicagoPizzaStore struct {
	BasicPizzaStore
}

func NewChicagoPizzaStore() *ChicagoPizzaStore {
	result := new(ChicagoPizzaStore)
	result.BasicPizzaStore.createPizza = result.createPizza
	return result
}

func (c *ChicagoPizzaStore) createPizza(pizzaType string) Pizza {
	ingredientFactory := ingredient.NewChicagoIngredientFactory()
	var pizza Pizza
	switch pizzaType {
	case "cheese":
		pizza =  NewCheesePizza(ingredientFactory)
		pizza.setName("Chicago Style Cheese Pizza")
	case "veggie":
		pizza =  NewVeggiePizza(ingredientFactory)
		pizza.setName("Chicago Style Veggie Pizza")
	case "clam":
		pizza =  NewClamPizza(ingredientFactory)
		pizza.setName("Chicago Style Clam Pizza")
	case "pepperoni":
		pizza =  NewPepperoniPizza(ingredientFactory)
		pizza.setName("Chicago Style Pepperoni Pizza")
	}
	return pizza
}
