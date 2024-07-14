package command

import "github.com/hankeyyh/head-first-design-pattern-go/chapter-6-command/vendor"

type CeilingFanOffCommand struct {
	ceilingFan *vendor.CeilingFan
	prevSpeed vendor.CeilingFanLevel
}

func NewCeilingFanOffCommand(ceilingFan *vendor.CeilingFan) *CeilingFanOffCommand {
	return &CeilingFanOffCommand{ceilingFan: ceilingFan}
}

func (c *CeilingFanOffCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.Off()
}

func (c *CeilingFanOffCommand) Undo() {
	switch c.prevSpeed {
	case vendor.High:
		c.ceilingFan.High()
	case vendor.Medium:
		c.ceilingFan.Medium()
	case vendor.Low:
		c.ceilingFan.Low()
	default:
		c.ceilingFan.Off()
	}
}

type CeilingFanHighCommand struct {
	ceilingFan *vendor.CeilingFan
	prevSpeed vendor.CeilingFanLevel
}

func NewCeilingFanHighCommand(ceilingFan *vendor.CeilingFan) *CeilingFanHighCommand {
	return &CeilingFanHighCommand{ceilingFan: ceilingFan}
}

func (c *CeilingFanHighCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.High()
}

func (c *CeilingFanHighCommand) Undo() {
	switch c.prevSpeed {
	case vendor.High:
		c.ceilingFan.High()
	case vendor.Medium:
		c.ceilingFan.Medium()
	case vendor.Low:
		c.ceilingFan.Low()
	default:
		c.ceilingFan.Off()
	}
}

type CeilingFanMediumCommand struct {
	ceilingFan *vendor.CeilingFan
	prevSpeed vendor.CeilingFanLevel
}

func NewCeilingFanMediumCommand(ceilingFan *vendor.CeilingFan) *CeilingFanMediumCommand {
	return &CeilingFanMediumCommand{ceilingFan: ceilingFan}
}

func (c *CeilingFanMediumCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.Medium()
}

func (c *CeilingFanMediumCommand) Undo() {
	switch c.prevSpeed {
	case vendor.High:
		c.ceilingFan.High()
	case vendor.Medium:
		c.ceilingFan.Medium()
	case vendor.Low:
		c.ceilingFan.Low()
	default:
		c.ceilingFan.Off()
	}
}