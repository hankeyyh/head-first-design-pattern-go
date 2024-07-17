package main

import "fmt"

type GumballMachine struct {
	soldState       State
	soldOutState    State
	noQuarterState  State
	hasQuarterState State
	winnerState     State

	curState State
	numGumballs int
}

func NewGumballMachine(cnt int) *GumballMachine {
	gm := new(GumballMachine)
	soldState := NewSoldState(gm)
	soldOutState := NewSoldOutState(gm)
	noQuarterState := NewNoQuarterState(gm)
	hasQuarterState := NewHasQuarterState(gm)
	winnerState := NewWinnerState(gm)

	gm.soldState = soldState
	gm.soldOutState = soldOutState
	gm.noQuarterState = noQuarterState
	gm.hasQuarterState = hasQuarterState
	gm.winnerState = winnerState
	gm.numGumballs = cnt	

	if cnt > 0 {
		gm.curState = noQuarterState
	} else {
		gm.curState = soldOutState
	}
	return gm
} 

func (gm *GumballMachine) InsertQuarter() {
	gm.curState.InsertQuarter()
}

func (gm *GumballMachine) EjectQuarter() {
	gm.curState.EjectQuarter()
}

func (gm *GumballMachine) TurnCrank() {
	gm.curState.TurnCrank()
	gm.curState.Dispense()
}

func (gm *GumballMachine) SetState(state State) {
	gm.curState = state
}

func (gm *GumballMachine) GetState() State {
	return gm.curState
}

func (gm *GumballMachine) GetSoldState() State {
	return gm.soldState
}

func (gm *GumballMachine) GetSoldOutState() State {
	return gm.soldOutState
}

func (gm *GumballMachine) GetNoQuarterState() State {
	return gm.noQuarterState
}

func (gm *GumballMachine) GetHasQuarterState() State {
	return gm.hasQuarterState
}

func (gm *GumballMachine) GetWinnerState() State {
	return gm.winnerState
}


func (gm *GumballMachine) ReleaseBall() {
	println("A gumball comes rolling out the slot...")
	if gm.numGumballs > 0 {
		gm.numGumballs--
	}
}

func (gm *GumballMachine) GetCount() int {
	return gm.numGumballs
}

func (gm *GumballMachine) ToString() string {
	return fmt.Sprintf("-----Mighty Gumball, Inc.-----\nInventory: %d gumballs\nMachine is %s\n", gm.numGumballs, gm.curState.ToString())
}