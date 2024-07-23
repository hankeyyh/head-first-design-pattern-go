package main

type DuckFactory interface {
	CreateMallardDuck() Quackable
	CreateRedHeadDuck() Quackable
	CreateDuckCall() Quackable
	CreateRubberDuck() Quackable
}

type CountingDuckFactory struct {

}

func NewCountingDuckFactory() *CountingDuckFactory {
	return &CountingDuckFactory{}
}

func (d *CountingDuckFactory) CreateMallardDuck() Quackable {
	return NewQuackCounter(NewMallarDuck())
}

func (d *CountingDuckFactory) CreateRedHeadDuck() Quackable {
	return NewQuackCounter(NewRedHeadDuck())
}

func (d *CountingDuckFactory) CreateDuckCall() Quackable {
	return NewQuackCounter(NewDuckCall())
}

func (d *CountingDuckFactory) CreateRubberDuck() Quackable {
	return NewQuackCounter(NewRubberDuck())
}

type GooseFactory interface {
	CreateGoose() Quackable
}

type CountingGooseFactory struct {

}

func NewCountingGooseFactory() *CountingGooseFactory {
	return &CountingGooseFactory{}
}

func (g *CountingGooseFactory) CreateGoose() Quackable {
	return NewQuackCounter(NewGooseAdapter(NewGoose()))
}