package main

import (
	"fmt"
	"math/rand"
)

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()

	Refill()
	ToString() string
}

type SoldState struct {
	gumballMachine *GumballMachine
}

func NewSoldState(gumballMachine *GumballMachine) *SoldState {
	return &SoldState{gumballMachine: gumballMachine}
}

func (s *SoldState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

func (s *SoldState) EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

func (s *SoldState) TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball!")
}

func (s *SoldState) Dispense() {
	s.gumballMachine.ReleaseBall()
	if s.gumballMachine.GetCount() > 0 {
		s.gumballMachine.SetState(s.gumballMachine.GetNoQuarterState())
	} else {
		fmt.Println("Oops, out of gumballs!")
		s.gumballMachine.SetState(s.gumballMachine.GetSoldOutState())
	}
}

func (s *SoldState) Refill() {
}

func (s *SoldState) ToString() string {
	return "dispensing a gumball"
}

type SoldOutState struct {
	gumballMachine *GumballMachine
}

func NewSoldOutState(gumballMachine *GumballMachine) *SoldOutState {
	return &SoldOutState{gumballMachine: gumballMachine}
}

func (s *SoldOutState) InsertQuarter() {
	fmt.Println("You can't insert a quarter, the machine is sold out")
}

func (s *SoldOutState) EjectQuarter() {
	fmt.Println("You can't eject, you haven't inserted a quarter yet")
}

func (s *SoldOutState) TurnCrank() {
	fmt.Println("You turned, but there are no gumballs")
}

func (s *SoldOutState) Dispense() {
	fmt.Println("No gumball dispensed")
}

func (s *SoldOutState) Refill() {
	s.gumballMachine.SetState(s.gumballMachine.GetNoQuarterState())
}

func (s *SoldOutState) ToString() string {
	return "sold out"
}


type NoQuarterState struct {
	gumballMachine *GumballMachine
}

func NewNoQuarterState(gumballMachine *GumballMachine) *NoQuarterState {
	return &NoQuarterState{gumballMachine: gumballMachine}
}

func (n *NoQuarterState) InsertQuarter() {
	fmt.Println("You inserted a quarter")
	n.gumballMachine.SetState(n.gumballMachine.GetHasQuarterState())
}

func (n *NoQuarterState) EjectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}

func (n *NoQuarterState) TurnCrank() {
	fmt.Println("You turned, but there's no quarter")
}

func (n *NoQuarterState) Dispense() {
	fmt.Println("You need to pay first")
}

func (n *NoQuarterState) Refill() {
	
}

func (n *NoQuarterState) ToString() string {
	return "waiting for quarter"
}


type HasQuarterState struct {
	gumballMachine *GumballMachine
}

func NewHasQuarterState(gumballMachine *GumballMachine) *HasQuarterState {
	return &HasQuarterState{gumballMachine: gumballMachine}
}

func (h *HasQuarterState) InsertQuarter() {
	fmt.Println("You can't insert another quarter")
}

func (h *HasQuarterState) EjectQuarter() {
	fmt.Println("Quarter returned")
	h.gumballMachine.SetState(h.gumballMachine.GetNoQuarterState())
}

func (h *HasQuarterState) TurnCrank() {
	fmt.Println("you turned...")
	if rand.Int() % 5 == 0 && h.gumballMachine.GetCount() > 1 {
		h.gumballMachine.SetState(h.gumballMachine.GetWinnerState())
	} else {
		h.gumballMachine.SetState(h.gumballMachine.GetSoldState())
	}
}

func (h *HasQuarterState) Dispense() {
	fmt.Println("No gumball dispensed")
}

func (h *HasQuarterState) Refill() {
	
}

func (h *HasQuarterState) ToString() string {
	return "waiting for turn of crank"
}

type WinnerState struct {
	gumballMachine *GumballMachine
}


func NewWinnerState(gumballMachine *GumballMachine) *WinnerState {
	return &WinnerState{gumballMachine: gumballMachine}
}

func (w *WinnerState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a Gumball")
}

func (w *WinnerState) EjectQuarter() {
	fmt.Println("Please wait, we're already giving you a Gumball")
}

func (w *WinnerState) TurnCrank() {
	fmt.Println("Turning again doesn't get you another gumball!")
}

func (w *WinnerState) Dispense() {
	w.gumballMachine.ReleaseBall()
	if w.gumballMachine.GetCount() == 0 {
		w.gumballMachine.SetState(w.gumballMachine.GetSoldOutState())
	} else {
		w.gumballMachine.ReleaseBall()
		fmt.Println("YOU'RE A WINNER! You got two gumballs for your quarter")
		if w.gumballMachine.GetCount() > 0 {
			w.gumballMachine.SetState(w.gumballMachine.GetNoQuarterState())
		} else {
			fmt.Println("Oops, out of gumballs!")
			w.gumballMachine.SetState(w.gumballMachine.GetSoldOutState())
		}
	}
}

func (w *WinnerState) Refill() {
	
}

func (w *WinnerState) ToString() string {
	return "despensing two gumballs for your quarter, because YOU'RE A WINNER!"
}