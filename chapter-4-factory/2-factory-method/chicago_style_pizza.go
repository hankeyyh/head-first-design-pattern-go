package main


type ChicagoStyleCheesePizza struct {
	BasicPizza
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	return &ChicagoStyleCheesePizza{
		BasicPizza: BasicPizza{
			name: "Chicago Style Deep Dish Cheese Pizza",
			dough: "Extra Thick Crust Dough",
			sauce: "Plum Tomato Sauce",
			toppings: []string{"Shredded Mozzarella Cheese"},
		},
	}
}

type ChicagoStyleVeggiePizza struct {
	BasicPizza
}

func NewChicagoStyleVeggiePizza() *ChicagoStyleVeggiePizza {
	return &ChicagoStyleVeggiePizza{
		BasicPizza: BasicPizza{
			name: "Chicago Style Deep Dish Veggie Pizza",
			dough: "Extra Thick Crust Dough",
			sauce: "Plum Tomato Sauce",
			toppings: []string{"Shredded Mozzarella Cheese", "Black Olives", "Spinach", "Eggplant"},
		},
	}
}

type ChicagoStyleClamPizza struct {
	BasicPizza
}

func NewChicagoStyleClamPizza() *ChicagoStyleClamPizza {
	return &ChicagoStyleClamPizza{
		BasicPizza: BasicPizza{
			name: "Chicago Style Clam Pizza",
			dough: "Extra Thick Crust Dough",
			sauce: "Plum Tomato Sauce",
			toppings: []string{"Shredded Mozzarella Cheese", "Frozen Clams from Chesapeake Bay"},
		},
	}
}
type ChicagoStylePepperoniPizza struct {
	BasicPizza
}

func NewChicagoStylePepperoniPizza() *ChicagoStylePepperoniPizza {
	return &ChicagoStylePepperoniPizza{
		BasicPizza: BasicPizza{
			name: "Chicago Style Pepperoni Pizza",
			dough: "Extra Thick Crust Dough",
			sauce: "Plum Tomato Sauce",
			toppings: []string{"Shredded Mozzarella Cheese", "Black Olives", "Spinach", "Eggplant", "Sliced Pepperoni"},
		},
	}
}