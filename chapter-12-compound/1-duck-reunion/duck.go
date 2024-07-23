package main

import (
	"fmt"

	"github.com/hankeyyh/head-first-design-pattern-go/chapter-12-compound/1-duck-reunion/observer"
)


type MallarDuck struct {
	observer.Observable
}

func NewMallarDuck() *MallarDuck {
	return &MallarDuck{}
}

func (m *MallarDuck) Quack() {
	fmt.Println("Quack")
}

type RedHeadDuck struct {

}

func NewRedHeadDuck() *RedHeadDuck {
	return &RedHeadDuck{}
}

func (r *RedHeadDuck) Quack() {
	fmt.Println("Quack")
}

type DuckCall struct {

}

func NewDuckCall() *DuckCall {
	return &DuckCall{}
}

func (d *DuckCall) Quack() {
	fmt.Println("Kwak")
}

type RubberDuck struct {

}

func NewRubberDuck() *RubberDuck {
	return &RubberDuck{}
}

func (r *RubberDuck) Quack() {
	fmt.Println("Squeak")
}