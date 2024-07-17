package main

import "fmt"

func main() {
	gumballMachine := NewGumballMachine(10)

	fmt.Println(gumballMachine.ToString())
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Println(gumballMachine.ToString())
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Println(gumballMachine.ToString())
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
}