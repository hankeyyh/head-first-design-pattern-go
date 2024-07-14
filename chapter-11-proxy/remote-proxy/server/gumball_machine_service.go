package main

import (
	"github.com/hankeyyh/head-first-design-pattern-go/chapter-11-proxy/remote-proxy/proto"
)

type GumballMachineService interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	GetCount() int
}

type GumballMachineServiceImp struct {
	machine *GumballMachine
}

func NewGumballMachineServiceImp(machine *GumballMachine) *GumballMachineServiceImp {
	return &GumballMachineServiceImp{machine: machine}
}


func (g *GumballMachineServiceImp) InsertQuarter(args *proto.EmptyRequest, resp *proto.EmptyResponse) error {
	g.machine.InsertQuarter()
	return nil
}

func (g *GumballMachineServiceImp) EjectQuarter(args *proto.EmptyRequest, resp *proto.EmptyResponse) error {
	g.machine.EjectQuarter()
	return nil
}

func (g *GumballMachineServiceImp) TurnCrank(args *proto.EmptyRequest, resp *proto.EmptyResponse) error {
	g.machine.TurnCrank()
	return nil
}

func (g *GumballMachineServiceImp) GetCount(args *proto.EmptyRequest, resp *proto.GetCountResponse) error {
	resp.Count = int32(g.machine.GetCount())
	return nil
}