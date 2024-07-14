package main

import (
	"fmt"
)

type Tuner struct {
	description string
	frequency float64
}

func NewTuner(description string) *Tuner {
	return &Tuner{description: description}
}

func (t *Tuner) On() {
	fmt.Println(t.description + " on")
}

func (t *Tuner) Off() {
	fmt.Println(t.description + " off")
}

func (t *Tuner) SetFrequency(frequency float64) {
	fmt.Println(t.description + " setting frequency to", frequency)
	t.frequency = frequency
}

func (t *Tuner) SetAm() {
	fmt.Println(t.description + " setting AM mode")
}

func (t *Tuner) SetFm() {
	fmt.Println(t.description + " setting FM mode")
}

func (t *Tuner) ToString() string {
	return t.description
}