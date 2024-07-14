package command

import "github.com/hankeyyh/head-first-design-pattern-go/chapter-6-command/vendor"

type HothubOffCommand struct {
	hothub *vendor.Hothub
}

func NewHothubOffCommand(hothub *vendor.Hothub) *HothubOffCommand {
	return &HothubOffCommand{hothub: hothub}
}

func (h *HothubOffCommand) Execute() {
	h.hothub.SetTemperature(98)
	h.hothub.Off()
}

func (h *HothubOffCommand) Undo() {
	h.hothub.On()
}

type HothubOnCommand struct {
	hothub *vendor.Hothub
}

func NewHothubOnCommand(hothub *vendor.Hothub) *HothubOnCommand {
	return &HothubOnCommand{hothub: hothub}
}

func (h *HothubOnCommand) Execute() {
	h.hothub.On()
	h.hothub.SetTemperature(104)
	h.hothub.Circulate()
}

func (h *HothubOnCommand) Undo() {
	h.hothub.Off()
}