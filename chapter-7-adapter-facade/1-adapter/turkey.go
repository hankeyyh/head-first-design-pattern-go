package main

import "fmt"

type Turkey interface {
	Gobble()
	Fly()
}

type WildTurkey struct {

}

func NewWildTurkey() *WildTurkey {
	return &WildTurkey{}
}

func (w *WildTurkey) Gobble() {
	fmt.Println("Gobble gobble")
}

func (w *WildTurkey) Fly() {
	fmt.Println("I'm flying a short distance")
}