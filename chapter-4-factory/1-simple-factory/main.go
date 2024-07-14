package main

/*
simple factory: compare to design pattern, it's more like an programming idiom.
Using a factory to avoid use 'new' keyword. 
*/
func main() {
	store := NewPizzaStore(NewSimplePizzaFactory())
	pizza := store.orderPizza("cheese")
	pizza.toString()
	pizza = store.orderPizza("veggie")
	pizza.toString()
	pizza = store.orderPizza("clam")
	pizza.toString()
	pizza = store.orderPizza("pepperoni")
	pizza.toString()
}