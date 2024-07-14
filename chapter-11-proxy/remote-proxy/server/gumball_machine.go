package main

type GumballMachine struct {
	count int

	noQuarterState State
	hasQuarterState State
	soldState State
	soldOutState State
	state State
}

func NewGumballMachine(count int) *GumballMachine {
	g := &GumballMachine{count: count}
	g.noQuarterState = NewNoQuarterState(g)
	g.hasQuarterState = NewHasQuarterState(g)
	g.soldState = NewSoldState(g)
	g.soldOutState = NewSoldOutState(g)
	if count > 0 {
		g.state = NewNoQuarterState(g)
	} else {
		g.state = NewSoldOutState(g)
	}
	return g
}

func (g *GumballMachine) GetCount() int {
	return g.count
}

func (g *GumballMachine) SetState(s State) {
	g.state = s
}

func (g *GumballMachine) InsertQuarter() {
	g.state.InsertQuarter()
}

func (g *GumballMachine) EjectQuarter() {
	g.state.EjectQuarter()
}

func (g *GumballMachine) TurnCrank() {
	g.state.TurnCrank()
	g.state.Dispense()
}

func (g *GumballMachine) ReleaseGumball() {
	if g.count <= 0 {
		return
	}
	println("A gumball comes rolling out the slot...")
	g.count--
}