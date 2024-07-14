package main

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
}

type NoQuarterState struct {
	machine *GumballMachine
}

func NewNoQuarterState(machine *GumballMachine) *NoQuarterState {
	return &NoQuarterState{machine: machine}
}

func (n *NoQuarterState) InsertQuarter() {
	println("You inserted a quarter")
	n.machine.SetState(n.machine.hasQuarterState)
}

func (n *NoQuarterState) EjectQuarter() {
	println("You haven't inserted a quarter")
}

func (n *NoQuarterState) TurnCrank() {
	println("You turned, but there's no quarter")
}

func (n *NoQuarterState) Dispense() {
	println("You need to pay first")
}

type HasQuarterState struct {
	machine *GumballMachine
}

func NewHasQuarterState(machine *GumballMachine) *HasQuarterState {
	return &HasQuarterState{machine: machine}
}

func (h *HasQuarterState) InsertQuarter() {
	println("You can't insert another quarter")
}

func (h *HasQuarterState) EjectQuarter() {
	println("Quarter returned")
	h.machine.SetState(h.machine.noQuarterState)
}

func (h *HasQuarterState) TurnCrank() {
	println("gumball is comming...")
	h.machine.SetState(h.machine.soldState)
}

func (h *HasQuarterState) Dispense() {
	println("No gumball dispensed")
}

type SoldState struct {
	machine *GumballMachine
}

func NewSoldState(machine *GumballMachine) *SoldState {
	return &SoldState{machine: machine}
}

func (s *SoldState) InsertQuarter() {
	println("Please wait, we're already giving you a gumball")
}

func (s *SoldState) EjectQuarter() {
	println("Sorry, you already turned the crank")
}

func (s *SoldState) TurnCrank() {
	println("Turning twice doesn't get you another gumball!")
}

func (s *SoldState) Dispense() {
	s.machine.ReleaseGumball()
	if s.machine.GetCount() > 0 {
		s.machine.SetState(s.machine.noQuarterState)
	} else {
		println("Oops, out of gumballs!")
		s.machine.SetState(s.machine.soldOutState)
	}
}

type SoldOutState struct {
	machine *GumballMachine
}

func NewSoldOutState(machine *GumballMachine) *SoldOutState {
	return &SoldOutState{machine: machine}
}

func (s *SoldOutState) InsertQuarter() {
	println("You can't insert a quarter, the machine is sold out")
}

func (s *SoldOutState) EjectQuarter() {
	println("You can't eject, you haven't inserted a quarter yet")
}

func (s *SoldOutState) TurnCrank() {
	println("You turned, but there are no gumballs")
}

func (s *SoldOutState) Dispense() {
	println("No gumball dispensed")
}
