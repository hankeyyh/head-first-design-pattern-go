package main

import "github.com/hankeyyh/head-first-design-pattern-go/chapter-12-compound/1-duck-reunion/observer"

type Flock struct {
	quackList []Quackable
}

func NewFlock() *Flock {
	return &Flock{}
}

func (f *Flock) Add(quack Quackable) {
	f.quackList = append(f.quackList, quack)
}

func (f *Flock) Quack() {
	for _, quack := range f.quackList {
		quack.Quack()
	}
}

// flock that can be ovserved
type FlockObservable struct {
	quackList []Quackable_QuackObservable
}

func NewFlockObservable() *FlockObservable {
	return &FlockObservable{}
}

func (f *FlockObservable) Add(quack Quackable_QuackObservable) {
	f.quackList = append(f.quackList, quack)
}

func (f *FlockObservable) Quack() {
	for _, quack := range f.quackList {
		quack.Quack()
	}
}

func (f *FlockObservable) RegisterObserver(observer observer.Observer) {
	for _, quack := range f.quackList {
		quack.RegisterObserver(observer)
	}
}

func (f *FlockObservable) NotifyObservers() {
	
}