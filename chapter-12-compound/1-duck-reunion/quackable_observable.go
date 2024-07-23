package main

import "github.com/hankeyyh/head-first-design-pattern-go/chapter-12-compound/1-duck-reunion/observer"

type Quackable_QuackObservable interface {
	Quackable
	observer.QuackObservable
}