package main

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