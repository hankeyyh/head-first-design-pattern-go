package main

type BeverageSize int

const (
	TALL BeverageSize = 1
	GRANDE BeverageSize = 2
	VENTI BeverageSize = 3
)

type BeverageSizeHandler struct {
	size BeverageSize
}

func (b *BeverageSizeHandler) SetSize(size BeverageSize) {
	b.size = size
}

func (b *BeverageSizeHandler) GetSize() BeverageSize {
	return b.size
}

type Beverage interface {
	GetDescription() string
	Cost() float64
	GetSize() BeverageSize
}

// HouseBlend is a concrete implementation of Beverage
type HouseBlend struct {
	BeverageSizeHandler
}

func NewHouseBlend() HouseBlend {
	return HouseBlend{}
}

func (h HouseBlend) GetDescription() string {
	return "House Blend Coffee"
}

func (h HouseBlend) Cost() float64 {
	return 0.89
}

// DarkRoast is a concrete implementation of Beverage
type DarkRoast struct {
	BeverageSizeHandler
}

func NewDarkRoast() DarkRoast {
	return DarkRoast{}
}

func (d DarkRoast) GetDescription() string {
	return "Dark Roast Coffee"
}

func (d DarkRoast) Cost() float64 {
	return 0.99
}

// Decaf is a concrete implementation of Beverage
type Decaf struct {

}

func NewDecaf() Decaf {
	return Decaf{}
}

func (d Decaf) GetDescription() string {
	return "Decaf Coffee"
}

func (d Decaf) Cost() float64 {
	return 1.05
}

// Espresso is a concrete implementation of Beverage
type Espresso struct {

}

func NewEspresso() Espresso {
	return Espresso{}
}

func (e Espresso) GetDescription() string {
	return "Espresso"
}

func (e Espresso) Cost() float64 {
	return 1.99
}