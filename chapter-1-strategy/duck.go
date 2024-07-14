package main

import "fmt"

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d *Duck) performFly() {
	d.flyBehavior.Fly()
}

func (d *Duck) performQuack() {
	d.quackBehavior.Quack()
}

func (d *Duck) swim() {
	fmt.Println("All ducks float, even decoys!")
}

func (d *Duck) display() {
	panic("Implement me")
}

func (d *Duck) SetFlyBehavior(fb FlyBehavior) {
	d.flyBehavior = fb
}

func (d *Duck) SetQuackBehavior(qb QuackBehavior) {
	d.quackBehavior = qb
}

// MallarDuck is a duck that can fly and quack
type MallarDuck struct {
	Duck
}

func NewMallarDuck() *MallarDuck {
	m := &MallarDuck{}
	m.flyBehavior = &FlyWithWings{}
	m.quackBehavior = &Quack{}
	return m
}

func (m *MallarDuck) display() {
	fmt.Println("I'm a real Mallard duck")
}

// ModelDuck is a duck that can't fly
type ModelDuck struct {
	Duck
}

func NewModelDuck() *ModelDuck {
	m := &ModelDuck{}
	m.flyBehavior = &FlyNoWay{}
	m.quackBehavior = &Quack{}
	return m
}

func (m *ModelDuck) display() {
	fmt.Println("I'm a model duck")
}


// DecoyDuck is a duck that can't fly and quack
type DecoyDuck struct {
	Duck
}

func NewDecoyDuck() *DecoyDuck {
	d := &DecoyDuck{}
	d.flyBehavior = &FlyNoWay{}
	d.quackBehavior = &MuteQuack{}
	return d
}

func (d *DecoyDuck) display() {
	fmt.Println("I'm a duck Decoy")
}

// RedHeadDuck is a duck that can fly and quack
type RedHeadDuck struct {
	Duck
}

func NewRedHeadDuck() *RedHeadDuck {
	r := &RedHeadDuck{}
	r.flyBehavior = &FlyWithWings{}
	r.quackBehavior = &Quack{}
	return r
}

func (r *RedHeadDuck) display() {
	fmt.Println("I'm a real Red Headed duck")
}

// RubberDuck is a duck that can't fly and squeak
type RubberDuck struct {
	Duck
}

func NewRubberDuck() *RubberDuck {
	r := &RubberDuck{}
	r.flyBehavior = &FlyNoWay{}
	r.quackBehavior = &Squeak{}
	return r
}

func (r *RubberDuck) display() {
	fmt.Println("I'm a rubber duckie")
}