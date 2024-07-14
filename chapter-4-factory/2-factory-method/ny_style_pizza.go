package main

type NYStyleCheesePizza struct {
	BasicPizza
}

func NewNewYorkStyleCheesePizza() *NYStyleCheesePizza {
	return &NYStyleCheesePizza{
		BasicPizza: BasicPizza{
			name: "NY Style Sauce and Cheese Pizza",
			dough: "Thin Crust Dough",
			sauce: "Marinara Sauce",
			toppings: []string{"Grated Reggiano Cheese"},
		},
	}
}

type NYStyleVeggiePizza struct {
	BasicPizza
}

func NewNewYorkStyleVeggiePizza() *NYStyleVeggiePizza {
	return &NYStyleVeggiePizza{
		BasicPizza: BasicPizza {
			name: "NY Style Veggie Pizza",
			dough: "Thin Crust Dough",
			sauce: "Marinara Sauce",
			toppings: []string{"Grated Reggiano Cheese", "Garlic", "Onion", "Mushrooms", "Red Pepper"},
		},
	}
}

type NYStyleClamPizza struct {
	BasicPizza
}

func NewNewYorkStyleClamPizza() *NYStyleClamPizza {
	return &NYStyleClamPizza{
		BasicPizza: BasicPizza {
			name: "NY Style Clam Pizza",
			dough: "Thin Crust Dough",
			sauce: "Marinara Sauce",
			toppings: []string{"Grated Reggiano Cheese", "Fresh Clams from Long Island Sound"},
		},
	}
}

type NYStylePepperoniPizza struct {
	BasicPizza
}

func NewNewYorkStylePepperoniPizza() *NYStylePepperoniPizza {
	return &NYStylePepperoniPizza{
		BasicPizza: BasicPizza {
			name: "NY Style Pepperoni Pizza",
			dough: "Thin Crust Dough",
			sauce: "Marinara Sauce",
			toppings: []string{"Grated Reggiano Cheese", "Sliced Pepperoni", "Garlic", "Onion", "Mushrooms", "Red Pepper"},
		},
	}
}
