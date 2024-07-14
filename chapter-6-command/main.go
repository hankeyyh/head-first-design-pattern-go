package main

import (
	"fmt"

	"github.com/hankeyyh/head-first-design-pattern-go/chapter-6-command/command"
	"github.com/hankeyyh/head-first-design-pattern-go/chapter-6-command/vendor"
)

func main() {
	livingRoomLight := vendor.NewLight("Living Room")
	hottub := vendor.NewHothub()
	bedroomFan := vendor.NewCeilingFan("Bed Room")

	lightOnCmd := command.NewLightOnCommand(livingRoomLight)
	lightOffCmd := command.NewLightOffCommand(livingRoomLight)
	hottubOnCmd := command.NewHothubOnCommand(hottub)
	hottubOffCmd := command.NewHothubOffCommand(hottub)
	ceilingFanHighCmd := command.NewCeilingFanHighCommand(bedroomFan)
	ceilingFanOffCmd := command.NewCeilingFanOffCommand(bedroomFan)

	partyOnCmd := command.NewMacroCommand([]command.Command{lightOnCmd, hottubOnCmd, ceilingFanHighCmd})
	partyOffCmd := command.NewMacroCommand([]command.Command{lightOffCmd, hottubOffCmd, ceilingFanOffCmd})

	remoteController := NewRemoteController()
	remoteController.SetCommand(0, lightOnCmd, lightOffCmd)
	remoteController.SetCommand(1, hottubOnCmd, hottubOffCmd)
	remoteController.SetCommand(2, ceilingFanHighCmd, ceilingFanOffCmd)
	remoteController.SetCommand(3, partyOnCmd, partyOffCmd)

	remoteController.ToString()
	remoteController.OnButtonWasPushed(0)
	remoteController.OffButtonWasPushed(0)
	remoteController.UndoButtonWasPushed()

	remoteController.OnButtonWasPushed(1)
	remoteController.OffButtonWasPushed(1)
	remoteController.UndoButtonWasPushed()

	remoteController.OnButtonWasPushed(2)
	remoteController.OffButtonWasPushed(2)
	remoteController.UndoButtonWasPushed()

	fmt.Println("------ Start Party Mode!! ------")
	remoteController.OnButtonWasPushed(3)
	remoteController.OffButtonWasPushed(3)
	remoteController.UndoButtonWasPushed()
}
