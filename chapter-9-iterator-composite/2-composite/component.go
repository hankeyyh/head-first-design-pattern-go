package main

type MenuComponment interface {
	Print()
	Add(child MenuComponment)
	Remove(child MenuComponment)
	GetChild(n int) MenuComponment
}
