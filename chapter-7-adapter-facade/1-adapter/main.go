package main

import "fmt"

func main() {
	duck := NewMalldDuck()
	duckAdapter := NewDuckAdapter(duck)

	turkey := NewWildTurkey()
	turkeyAdapter := NewTurkeyAdapter(turkey)

	fmt.Println("\nThe duck says...")
	testDuck(duck)

	fmt.Println("The turkey says...")
	testTurkey(turkey)

	fmt.Println("\nThe turkey adapter says...")
	testDuck(turkeyAdapter)

	fmt.Println("\nThe duck adapter says...")
	testTurkey(duckAdapter)
}

func testDuck(duck Duck) {
	duck.Quack()
	duck.Fly()
}

func testTurkey(turkey Turkey) {
	turkey.Gobble()
	turkey.Fly()
}