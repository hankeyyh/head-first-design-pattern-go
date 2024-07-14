package ingredient

type Veggies interface {
	ToVeggiesString() string
}

type BlackOlives struct {

}

func NewBlackOlives() *BlackOlives {
	return &BlackOlives{}
}

func (b *BlackOlives) ToVeggiesString() string {
	return "Black Olives"
}

type Eggplant struct {

}

func NewEggplant() *Eggplant {
	return &Eggplant{}
}

func (b *Eggplant) ToVeggiesString() string {
	return "Eggplant"
}

type Garlic struct {

}

func NewGarlic() *Garlic {
	return &Garlic{}
}

func (b *Garlic) ToVeggiesString() string {
	return "Garlic"
}

type Mushroom struct {

}

func NewMushroom() *Mushroom {
	return &Mushroom{}
}

func (b *Mushroom) ToVeggiesString() string {
	return "Mushroom"
}

type Onion struct {

}

func NewOnion() *Onion {
	return &Onion{}
}

func (b *Onion) ToVeggiesString() string {
	return "Onion"
}

type RedPepper struct {

}

func NewRedPepper() *RedPepper {
	return &RedPepper{}
}

func (b *RedPepper) ToVeggiesString() string {
	return "Red Pepper"
}

type Spinach struct {

}

func NewSpinach() *Spinach {
	return &Spinach{}
}

func (b *Spinach) ToVeggiesString() string {
	return "Spinach"
}