package main

import "fmt"

type Goose struct {

}

func NewGoose() *Goose {
	return &Goose{}
}

func (g *Goose) Honk() {
	fmt.Println("Honk")
}

type GooseAdapter struct {
	goose *Goose
}

func NewGooseAdapter(goose *Goose) *GooseAdapter {
	return &GooseAdapter{goose: goose}
}

func (g *GooseAdapter) Quack() {
	g.goose.Honk()
}