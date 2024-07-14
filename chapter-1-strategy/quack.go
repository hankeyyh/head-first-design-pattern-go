package main

import "fmt"

type QuackBehavior interface {
	Quack()
}

type Quack struct {

}

func (q Quack) Quack() {
	fmt.Println("Quack")
}

type MuteQuack struct {

}

func (q MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}

type FakeQuack struct {

}

func (q FakeQuack) Quack() {
	fmt.Println("Qwak")
}

type Squeak struct {

}

func (q Squeak) Quack() {
	fmt.Println("Squeak")
}