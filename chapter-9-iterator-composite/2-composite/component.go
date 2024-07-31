package main

type MenuComponment interface {
	GetName() string
	GetDescription() string
	GetPrice() float64
	IsVegetarian() bool
	Print()
	Add(child MenuComponment)
	Remove(child MenuComponment)
	GetChild(n int) MenuComponment
}
