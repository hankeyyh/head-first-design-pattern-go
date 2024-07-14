package main

/*
factory-method: let the factory subclass decide how to create concrete product.
compare to simple factory, we control the orderPizza procedure.
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