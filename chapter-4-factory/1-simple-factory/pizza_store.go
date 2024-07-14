package main

type PizzaStore struct {
	factory *SimplePizzaFactory
}

func NewPizzaStore(factory *SimplePizzaFactory) *PizzaStore {
	return &PizzaStore{factory: factory}
}

func (p *PizzaStore) orderPizza(pizzaType string) Pizza {
	pizza := p.factory.CreatePizza(pizzaType)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
	return pizza
}