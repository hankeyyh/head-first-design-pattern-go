package main

import "fmt"

type CaffeineBeverage interface {
	Brew()
	AddCondiments()
}

type Coffee struct {

}

func NewCoffee() *Coffee {
	return &Coffee{}
}

func (c *Coffee) Brew() {
	fmt.Println("ripping Coffee through filter")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding Sugar and Milk")
}


type Tea struct {

}

func NewTea() *Tea {
	return &Tea{}
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding Lemon")
}