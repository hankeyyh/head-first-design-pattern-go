package main

import "fmt"

/*
because there isn't inheritance in go.
it's bit hard to implement template method pattern as the book says.
it's implemented more like a strategy pattern.
*/
func main() {
	fmt.Println("----- Coffee prepareing -----")
	recipe1 := NewRecipe(NewCoffee())
	recipe1.prepareCaffineBeverage()
	
	fmt.Println("----- Tea prepareing -----")
	recipe2 := NewRecipe(NewTea())
	recipe2.prepareCaffineBeverage()
}
