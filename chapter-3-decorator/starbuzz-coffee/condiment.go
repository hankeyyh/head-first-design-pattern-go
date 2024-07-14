package main

type Condiment struct {
	beverage Beverage
}

func (c Condiment) GetSize() BeverageSize {
	return c.beverage.GetSize()
}


// Milk is a concrete implementation of Condiment
type Milk struct {
	Condiment
}

func NewMild(beverage Beverage) Beverage {
	return Milk{Condiment{beverage}}
}

func (m Milk) GetDescription() string {
	return m.beverage.GetDescription() + ", Milk"
}

func (m Milk) Cost() float64 {
	return 0.10 + m.beverage.Cost()
}

// Mocha is a concrete implementation of Condiment
type Mocha struct {
	Condiment
}

func NewMocha(beverage Beverage) Beverage {
	return Mocha{Condiment{beverage}}
}

func (m Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", Mocha"
}

func (m Mocha) Cost() float64 {
	return 0.20 + m.beverage.Cost()
}

// Soy is a concrete implementation of Condiment
type Soy struct {
	Condiment
}

func NewSoy(beverage Beverage) Beverage {
	return Soy{Condiment{beverage}}
}

func (s Soy) GetDescription() string {
	return s.beverage.GetDescription() + ", Soy"
}

func (s Soy) Cost() float64 {
	cost := s.beverage.Cost()
	switch s.beverage.GetSize() {
	case TALL:
		cost += 0.05
	case GRANDE:
		cost += 0.40
	case VENTI:
		cost += 0.80
	}
	return cost
}

// Whip is a concrete implementation of Condiment
type Whip struct {
	Condiment
}

func NewWhip(beverage Beverage) Beverage {
	return Whip{Condiment{beverage}}
}

func (w Whip) GetDescription() string {
	return w.beverage.GetDescription() + ", Whip"
}

func (w Whip) Cost() float64 {
	return 0.10 + w.beverage.Cost()
}