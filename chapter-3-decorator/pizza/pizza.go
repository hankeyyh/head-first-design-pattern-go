package main

type Pizza interface {
	GetDescription() string
	Cost() float64
}

// ThinCrustPizza is a concrete implementation of Pizza
type ThickcrustPizza struct {

}

func NewThickcrustPizza() ThickcrustPizza {
	return ThickcrustPizza{}
}

func (t ThickcrustPizza) GetDescription() string {
	return "ThickcrustPizza"
}

func (t ThickcrustPizza) Cost() float64 {
	return 10.0
}

// ThinCrustPizza is a concrete implementation of Pizza
type ThincrustPizza struct {

}

func NewThincrustPizza() ThincrustPizza {
	return ThincrustPizza{}
}

func (t ThincrustPizza) GetDescription() string {
	return "ThincrustPizza"
}

func (t ThincrustPizza) Cost() float64 {
	return 5.0
}