package main

import (
	"fmt"

	"github.com/hankeyyh/head-first-design-pattern-go/chapter-12-compound/1-duck-reunion/observer"
)

type MallarDuck struct {
	*observer.BasicObservable // make duck observable
}

func NewMallarDuck() *MallarDuck {
	result := &MallarDuck{}
	result.BasicObservable = &observer.BasicObservable{Duck: result}
	return result
}

func (m *MallarDuck) Quack() {
	fmt.Println("Quack")
	m.NotifyObservers()
}

type RedHeadDuck struct {
	*observer.BasicObservable
}

func NewRedHeadDuck() *RedHeadDuck {
	result := &RedHeadDuck{}
	result.BasicObservable = &observer.BasicObservable{Duck: result}
	return result
}

func (r *RedHeadDuck) Quack() {
	fmt.Println("Quack")
	r.NotifyObservers()
}

type DuckCall struct {
	*observer.BasicObservable
}

func NewDuckCall() *DuckCall {
	result := &DuckCall{}
	result.BasicObservable = &observer.BasicObservable{Duck: result}
	return result
}

func (d *DuckCall) Quack() {
	fmt.Println("Kwak")
	d.NotifyObservers()
}

type RubberDuck struct {
	*observer.BasicObservable
}

func NewRubberDuck() *RubberDuck {
	result := &RubberDuck{}
	result.BasicObservable = &observer.BasicObservable{Duck: result}
	return result
}

func (r *RubberDuck) Quack() {
	fmt.Println("Squeak")
	r.NotifyObservers()
}
