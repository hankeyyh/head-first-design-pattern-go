package main

import "fmt"

type MenuItem struct {
	name        string
	description string
	vegetarian  bool
	price       float64
}

func NewMenuItem(name string, desc string, veg bool, price float64) *MenuItem {
	return &MenuItem{name, desc, veg, price}
}

func (m *MenuItem) GetName() string {
	return m.name
}

func (m *MenuItem) GetDescription() string {
	return m.description
}

func (m *MenuItem) GetPrice() float64 {
	return m.price
}

func (m *MenuItem) IsVegetarian() bool {
	return m.vegetarian
}

func (m *MenuItem) Print() {
	fmt.Printf("%s-%.2f-%s\n", m.GetName(), m.GetPrice(), m.GetDescription())
}

func (m *MenuItem) Add(child MenuComponment) {

}

func (m *MenuItem) Remove(child MenuComponment) {

}

func (m *MenuItem) GetChild(n int) MenuComponment {
	return nil
}
