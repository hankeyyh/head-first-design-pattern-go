package main

func main() {
	gumballMachine := NewGumballMachineStub("localhost:1234")
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	gumballMachine.InsertQuarter()
	gumballMachine.EjectQuarter()

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	println("Left Gumball count:", gumballMachine.GetCount())
	println("Over")
}