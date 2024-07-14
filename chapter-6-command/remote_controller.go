package main

import (
	"fmt"

	"github.com/hankeyyh/head-first-design-pattern-go/chapter-6-command/command"
)

type RemoteController struct {
	onCommands []command.Command
	offCommands []command.Command
	undoCommand command.Command
}

func NewRemoteController() *RemoteController {
	remoteController := &RemoteController{}
	remoteController.onCommands = make([]command.Command, 7)
	remoteController.offCommands = make([]command.Command, 7)

	noCommand := command.NewNoCommand()
	for i := 0; i < 7; i++ {
		remoteController.onCommands[i] = noCommand
		remoteController.offCommands[i] = noCommand
	}
	remoteController.undoCommand = noCommand
	return remoteController
}

func (r *RemoteController) SetCommand(slot int, onCommand command.Command, offCommand command.Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

func (r *RemoteController) OnButtonWasPushed(slot int) {
	r.onCommands[slot].Execute()
	r.undoCommand = r.onCommands[slot]
}

func (r *RemoteController) OffButtonWasPushed(slot int) {
	r.offCommands[slot].Execute()
	r.undoCommand = r.offCommands[slot]
}

func (r *RemoteController) UndoButtonWasPushed() {
	r.undoCommand.Undo()
}

func (r *RemoteController) ToString() {
	fmt.Println("------ Remote Controller ------")
	for i := 0; i < len(r.onCommands); i++ {
		fmt.Println("[slot " + string(i) + "] ", r.onCommands[i], "    ", r.offCommands[i])
	}
	fmt.Println("[undo] ", r.undoCommand)
}