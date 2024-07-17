package main

type Recipe struct {
	beverage CaffeineBeverage
}

func NewRecipe(beverage CaffeineBeverage) *Recipe {
	return &Recipe{beverage: beverage}
}

func (r *Recipe) prepareCaffineBeverage() {
	r.BoilWater()
	r.beverage.Brew()
	r.PourInCup()
	r.beverage.AddCondiments()
}

func (r *Recipe) BoilWater() {
	println("Boiling water")
}

func (r *Recipe) PourInCup() {
	println("Pouring into cup")
}