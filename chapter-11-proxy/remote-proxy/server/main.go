package main

import (
	"net"
	"net/rpc"

	"github.com/mars9/codec"
)

func main() {
	ln, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 1234})
	if err != nil {
		panic(err)
	}
	svc := NewGumballMachineServiceImp(NewGumballMachine(5))
	rpc.RegisterName("GumballMachine", svc)
	println("Start server")
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		println("Accept connection")
		go rpc.ServeCodec(codec.NewServerCodec(conn))
	}
}
