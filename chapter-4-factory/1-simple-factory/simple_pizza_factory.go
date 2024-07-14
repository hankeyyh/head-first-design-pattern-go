package main

type SimplePizzaFactory struct {

}

func NewSimplePizzaFactory() *SimplePizzaFactory {
	return &SimplePizzaFactory{}
}

func (s *SimplePizzaFactory) CreatePizza(pizzaType string) Pizza {
	var pizza Pizza
	switch pizzaType {
	case "cheese":
		pizza = NewCheesePizza()
	case "pepperoni":
		pizza = NewPepperoniPizza()
	case "clam":
		pizza = NewClamPizza()
	case "veggie":
		pizza = NewVeggiePizza()
	}
	return pizza
}