package main

/*
abstract factory: create a family of products. 
compare to factory-method, we control the dough, sauce, toppings creation.
no need to define pizza for each style(NYStyleCheesePizza, ChicagoStyleCheesePizza).
just use style corresponded ingredients.
*/
func main() {
	nyStore := NewNewYorkPizzaStore()
	pizza := nyStore.orderPizza("cheese")
	pizza.toString()
	pizza = nyStore.orderPizza("veggie")
	pizza.toString()
	pizza = nyStore.orderPizza("clam")
	pizza.toString()
	pizza = nyStore.orderPizza("pepperoni")
	pizza.toString()

	chicagoStore := NewChicagoPizzaStore()
	pizza = chicagoStore.orderPizza("cheese")
	pizza.toString()
	pizza = chicagoStore.orderPizza("veggie")
	pizza.toString()
	pizza = chicagoStore.orderPizza("clam")
	pizza.toString()
	pizza = chicagoStore.orderPizza("pepperoni")
	pizza.toString()
}