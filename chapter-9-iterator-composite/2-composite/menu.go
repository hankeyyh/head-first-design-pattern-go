package main

import (
	"fmt"
)

type Menu interface {
	createIterator() Iterator
}

const DinerMenuMaxItems = 6

type BasicMenu struct {
	name        string
	description string
	menuComponments []MenuComponment
}

func NewBasicMenu(name, description string) *BasicMenu {
	return &BasicMenu{name, description, make([]MenuComponment, 0)}
}

func (b *BasicMenu) GetName() string {
	return b.name
}

func (b *BasicMenu) GetDescription() string {
	return b.description
}

func (b *BasicMenu) GetPrice() float64 {
	return 0
}

func (b *BasicMenu) IsVegetarian() bool {
	return false
}

func (b *BasicMenu) Print() {
	fmt.Printf("\n%s, %s\n", b.GetName(), b.GetDescription())
	fmt.Println("---------------------")

	iterator := b.createIterator()
	for iterator.HasNext() {
		iterator.Next().Print()
	}
}

func (b *BasicMenu) Add(child MenuComponment) {
	b.menuComponments = append(b.menuComponments, child)
}

func (b *BasicMenu) Remove(child MenuComponment) {
	for i, v := range b.menuComponments {
		if v == child {
			b.menuComponments = append(b.menuComponments[:i], b.menuComponments[i+1:]...)
			break
		}
	}
}

func (b *BasicMenu) GetChild(n int) MenuComponment {
	return b.menuComponments[n]
}

func (b *BasicMenu) createIterator() Iterator {
	return NewBasicMenuIterator(b.menuComponments)
}