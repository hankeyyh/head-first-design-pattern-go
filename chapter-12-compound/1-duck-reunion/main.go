package main

import (
	"fmt"

	"github.com/hankeyyh/head-first-design-pattern-go/chapter-12-compound/1-duck-reunion/observer"
)


func DuckSimulator() {
	// Version 1
	var mallardDuck, redHeadDuck, duckCall, rubberDuck, gooseDuck Quackable

	mallardDuck = NewMallarDuck()
	redHeadDuck = NewRedHeadDuck()
	duckCall = NewDuckCall()
	rubberDuck = NewRubberDuck()
	gooseDuck = NewGooseAdapter(NewGoose())

	fmt.Println("Duck Simulator1: with goose adapter")
	simulate(mallardDuck, redHeadDuck, duckCall, rubberDuck, gooseDuck)

	// Version 2
	mallardDuck = NewQuackCounter(mallardDuck)
	redHeadDuck = NewQuackCounter(redHeadDuck)
	duckCall = NewQuackCounter(duckCall)
	rubberDuck = NewQuackCounter(rubberDuck)
	gooseDuck = NewQuackCounter(gooseDuck)

	fmt.Println("\nDuck Simulator2: with decorator")
	simulate(mallardDuck, redHeadDuck, duckCall, rubberDuck, gooseDuck)
	fmt.Println("The ducks quacked", numberOfQuacks, "times")

	// Version 3
	duckFactory := NewCountingDuckFactory()
	gooseFactory := NewCountingGooseFactory()
	mallardDuck = duckFactory.CreateMallardDuck()
	redHeadDuck = duckFactory.CreateRedHeadDuck()
	duckCall = duckFactory.CreateDuckCall()
	rubberDuck = duckFactory.CreateRubberDuck()
	gooseDuck = gooseFactory.CreateGoose()

	fmt.Println("\nDuck Simulator3: with abstract factory")
	simulate(mallardDuck, redHeadDuck, duckCall, rubberDuck, gooseDuck)
	fmt.Println("The ducks quacked", numberOfQuacks, "times")

	// Version 4
	duckFlock := NewFlock()
	duckFlock.Add(redHeadDuck)
	duckFlock.Add(duckCall)
	duckFlock.Add(rubberDuck)
	duckFlock.Add(gooseDuck)

	mallardFlock := NewFlock()
	mallardFlock.Add(mallardDuck)
	mallardFlock.Add(duckFactory.CreateMallardDuck())
	mallardFlock.Add(duckFactory.CreateMallardDuck())
	mallardFlock.Add(duckFactory.CreateMallardDuck())

	duckFlock.Add(mallardFlock)

	fmt.Println("\nDuck Simulator4: with composite - Whole Flock simulation")
	simulate(duckFlock)

	fmt.Println("\nDuck Simulator: with composite - Mallard Flock simulation")
	simulate(mallardFlock)

	fmt.Println("The ducks quacked", numberOfQuacks, "times")

	// Version 5
	duckFlockObservable := NewFlockObservable()
	duckFlockObservable.Add(NewMallarDuck())
	duckFlockObservable.Add(NewRedHeadDuck())
	duckFlockObservable.Add(NewDuckCall())
	duckFlockObservable.Add(NewRubberDuck())
	
	duckFlockObservable.RegisterObserver(observer.NewQuackologist())

	fmt.Println("\nDuck Simulator5: With Observer")
	simulate(duckFlockObservable)
}

func simulate(quackList ...Quackable) {
	for _, quack := range quackList {
		quack.Quack()
	}
}

func main() {
	DuckSimulator()
}