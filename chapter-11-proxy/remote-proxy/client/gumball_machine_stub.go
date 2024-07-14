package main

import (
	"net"
	"net/rpc"

	"github.com/mars9/codec"

	"github.com/hankeyyh/head-first-design-pattern-go/chapter-11-proxy/remote-proxy/proto"
)

type GumballMachineService interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	GetCount() int
}

type GumballMachineStub struct {
	client *rpc.Client
}

func NewGumballMachineStub(addr string) *GumballMachineStub {
	conn, err := net.Dial("tcp", addr)
	client := rpc.NewClientWithCodec(codec.NewClientCodec(conn))
	if err != nil {
		panic(err)
	}
	return &GumballMachineStub{
		client: client,
	}
}

func (g *GumballMachineStub) InsertQuarter() {
	err := g.client.Call("GumballMachine.InsertQuarter", &proto.EmptyRequest{}, &proto.EmptyResponse{})
	if err != nil {
		panic(err)
	}
}

func (g *GumballMachineStub) EjectQuarter() {
	g.client.Call("GumballMachine.EjectQuarter", &proto.EmptyRequest{}, &proto.EmptyResponse{})
}

func (g *GumballMachineStub) TurnCrank() {
	g.client.Call("GumballMachine.TurnCrank", &proto.EmptyRequest{}, &proto.EmptyResponse{})
}

func (g *GumballMachineStub) GetCount() int {
	resp := &proto.GetCountResponse{}
	g.client.Call("GumballMachine.GetCount", &proto.EmptyRequest{}, resp)
	return int(resp.GetCount())
}
