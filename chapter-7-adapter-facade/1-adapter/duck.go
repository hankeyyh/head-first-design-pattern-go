package main

import "fmt"

type Duck interface {
	Quack()
	Fly()
}

type MallardDuck struct {
	
}

func NewMalldDuck() *MallardDuck {
	return &MallardDuck{}
}

func (m *MallardDuck) Quack() {
	fmt.Println("Quack")
}

func (m *MallardDuck) Fly() {
	fmt.Println("I'm flying")
}