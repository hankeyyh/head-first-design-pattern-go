package ingredient

type IngredientFactory interface {
	CreateCheese() Cheese
	CreateClams() Clams
	CreateDough() Dough
	CreatePepperoni() Pepperoni
	CreateSauce() Sauce
	CreateVeggies() []Veggies
}

type NYIngredientFactory struct {

}

func NewNYIngredientFactory() *NYIngredientFactory {
	return &NYIngredientFactory{}
}

func (n *NYIngredientFactory) CreateCheese() Cheese {
	return NewReggianoCheese()
}

func (n *NYIngredientFactory) CreateClams() Clams {
	return NewFreshClams()
}

func (n *NYIngredientFactory) CreateDough() Dough {
	return NewThinCrustDough()
}

func (n *NYIngredientFactory) CreatePepperoni() Pepperoni {
	return NewSlicedPepperoni()
}

func (n *NYIngredientFactory) CreateSauce() Sauce {
	return NewMarinaraSauce()
}

func (n *NYIngredientFactory) CreateVeggies() []Veggies {
	return []Veggies{NewGarlic(), NewOnion(), NewMushroom(), NewRedPepper()}
}


type ChicagoIngredientFactory struct {

}

func NewChicagoIngredientFactory() *ChicagoIngredientFactory {
	return &ChicagoIngredientFactory{}
}

func (c *ChicagoIngredientFactory) CreateCheese() Cheese {
	return NewMozzarellaCheese()
}

func (c *ChicagoIngredientFactory) CreateClams() Clams {
	return NewFrozenClams()
}

func (c *ChicagoIngredientFactory) CreateDough() Dough {
	return NewThickCrustDough()
}

func (c *ChicagoIngredientFactory) CreatePepperoni() Pepperoni {
	return NewSlicedPepperoni()
}

func (c *ChicagoIngredientFactory) CreateSauce() Sauce {
	return NewPlumTomatoSauce()
}

func (c *ChicagoIngredientFactory) CreateVeggies() []Veggies {
	return []Veggies{NewBlackOlives(), NewSpinach(), NewEggplant()}
}


