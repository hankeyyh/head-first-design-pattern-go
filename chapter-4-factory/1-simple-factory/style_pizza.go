package main

type CheesePizza struct {
	BasicPizza
}

func NewCheesePizza() *CheesePizza {
	return &CheesePizza{
		BasicPizza: BasicPizza{
			name: "Cheese Pizza",
			dough: "Regular Crust",
			sauce: "Marinara Pizza Sauce",
			toppings: []string{"Fresh Mozzarella", "Parmesan"},
		},
	}
}

type ClamPizza struct {
	BasicPizza
}

func NewClamPizza() *ClamPizza {
	return &ClamPizza{
		BasicPizza: BasicPizza{
			name: "Clam Pizza",
			dough: "Thin Crust",
			sauce: "White Garlic Sauce",
			toppings: []string{"Clams", "Grated Parmesan"},
		},
	}
}

type PepperoniPizza struct {
	BasicPizza
}

func NewPepperoniPizza() *PepperoniPizza {
	return &PepperoniPizza{
		BasicPizza: BasicPizza{
			name: "Pepperoni Pizza",
			dough: "Crust",
			sauce: "Marinara sauce",
			toppings: []string{"Sliced Pepperoni", "Sliced Onion", "Grated parmesan cheese"},
		},
	}
}

type VeggiePizza struct {
	BasicPizza
}

func NewVeggiePizza() *VeggiePizza {
	return &VeggiePizza{
		BasicPizza: BasicPizza{
			name: "Veggie Pizza",
			dough: "Crust",
			sauce: "Marinara sauce",
			toppings: []string{"Shredded mozzarella", "Grated parmesan", "Diced onion", "Sliced mushrooms", "Sliced red pepper", "Sliced black olives"},
		},
	}
}