package main

import "fmt"

type PizzaStore interface {
	createPizza(pizzaType string) Pizza
	orderPizza(pizzaType string) Pizza
}

type BasicPizzaStore struct {
	createPizza func(pizzaType string) Pizza
}

// for quality control, stick the orderPizza procedure. StyleStores provide their own creation method
func (d *BasicPizzaStore) orderPizza(pizzaType string) Pizza {
	pizza := d.createPizza(pizzaType)
	fmt.Println("--- Making a ", pizza.getName(), " ---")
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
	if pizzaType == "cheese" {
		return NewNewYorkStyleCheesePizza()
	} else if pizzaType == "veggie" {
		return NewNewYorkStyleVeggiePizza()
	} else if pizzaType == "clam" {
		return NewNewYorkStyleClamPizza()
	} else if pizzaType == "pepperoni" {
		return NewNewYorkStylePepperoniPizza()
	}
	return nil
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
	if pizzaType == "cheese" {
		return NewChicagoStyleCheesePizza()
	} else if pizzaType == "veggie" {
		return NewChicagoStyleVeggiePizza()
	} else if pizzaType == "clam" {
		return NewChicagoStyleClamPizza()
	} else if pizzaType == "pepperoni" {
		return NewChicagoStylePepperoniPizza()
	}
	return nil
}
