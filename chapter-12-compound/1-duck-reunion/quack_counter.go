package main

var numberOfQuacks int

type QuackCounter struct {
	quack Quackable
}

func NewQuackCounter(quack Quackable) *QuackCounter {
	return &QuackCounter{quack: quack}
}

func (q *QuackCounter) Quack() {
	numberOfQuacks++
	q.quack.Quack()
}

