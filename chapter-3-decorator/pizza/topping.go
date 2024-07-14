package main

type Cheese struct {
	pizza Pizza
}

func NewCheese(pizza Pizza) Pizza {
	return Cheese{pizza}
}

func (c Cheese) GetDescription() string {
	return c.pizza.GetDescription() + ", Cheese"
}

func (c Cheese) Cost() float64 {
	return 0.50 + c.pizza.Cost()
}

type Olives struct {
	pizza Pizza
}

func NewOlives(pizza Pizza) Pizza {
	return Olives{pizza}
}

func (o Olives) GetDescription() string {
	return o.pizza.GetDescription() + ", Olives"
}

func (o Olives) Cost() float64 {
	return 0.30 + o.pizza.Cost()
}