package command

import "github.com/hankeyyh/head-first-design-pattern-go/chapter-6-command/vendor"

type LightOnCommand struct {
	light *vendor.Light
}

func NewLightOnCommand(light *vendor.Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (l *LightOnCommand) Execute() {
	l.light.TurnOn()
}

func (l *LightOnCommand) Undo() {
	l.light.TurnOff()
}


type LightOffCommand struct {
	light *vendor.Light
}

func NewLightOffCommand(light *vendor.Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (l *LightOffCommand) Execute() {
	l.light.TurnOff()
}

func (l *LightOffCommand) Undo() {
	l.light.TurnOn()
}
