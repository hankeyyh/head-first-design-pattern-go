package main

import "fmt"

type Waitress struct {
	pancakeHouseMenu Menu
	dinerMenu Menu
	cafeMenu Menu
}

func NewWaitress(phm, dm, cm Menu) *Waitress {
	return &Waitress{phm, dm, cm}
}

func (w *Waitress) printMenu() {
	fmt.Printf("Menu\n---\nBreakfast\n")
	w.printSingleMenu(w.pancakeHouseMenu.createIterator())
	fmt.Printf("\nLunch\n")
	w.printSingleMenu(w.dinerMenu.createIterator())
	fmt.Printf("\nDinner\n")
	w.printSingleMenu(w.cafeMenu.createIterator())
}

func (w *Waitress) printSingleMenu(iterator Iterator) {
	for iterator.HasNext() {
		menuItem := iterator.Next()
		fmt.Printf("%s-%.2f-%s\n", menuItem.GetName(), menuItem.GetPrice(), menuItem.GetDescription())
	}
}

func (w *Waitress) printVegetarianMenu() {
	fmt.Printf("\nVegetarian Menu\n---\nBreakfast\n")
	w.printSingleVegetarianMenu(w.pancakeHouseMenu.createIterator())
	fmt.Printf("\nLunch\n")
	w.printSingleVegetarianMenu(w.dinerMenu.createIterator())
}

func (w *Waitress) printSingleVegetarianMenu(iterator Iterator) {
	for iterator.HasNext() {
		menuItem := iterator.Next()
		if menuItem.IsVegetarian() {
			fmt.Printf("%s-%.2f-%s\n", menuItem.GetName(), menuItem.GetPrice(), menuItem.GetDescription())
		}
	}
}