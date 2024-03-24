package main

import "fmt"

func main() {
	remoteControl := NewRemoteControl()

	livingRoomLight := NewLight("Living Room")

	livingRoomLightOn := NewLightOnCommand(*livingRoomLight)
	livingRoomLightOff := NewLightOffCommand(*livingRoomLight)

	remoteControl.SetCommand(0, livingRoomLightOn, livingRoomLightOff)

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	fmt.Println(remoteControl)
	remoteControl.UndoButtonWasPressed()
	remoteControl.OffButtonWasPushed(0)
	remoteControl.OnButtonWasPushed(0)
	fmt.Println(remoteControl)
	remoteControl.UndoButtonWasPressed()

	ceilingFan := NewCeilingFan("Living Room")
	ceilingFanMediumCommand := NewCeilingFanMediumCommand(*ceilingFan)
	ceilingFanHighCommand := NewCeilingFanHighCommand(*ceilingFan)
	ceilingFanOfffCommand := NewCeilingFanOffCommand(*ceilingFan)

	remoteControl.SetCommand(0, ceilingFanMediumCommand, ceilingFanOfffCommand)
	remoteControl.SetCommand(1, ceilingFanHighCommand, ceilingFanOfffCommand)

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	fmt.Println(remoteControl)
	remoteControl.UndoButtonWasPressed()

	remoteControl.OnButtonWasPushed(1)
	fmt.Println(remoteControl)
	remoteControl.UndoButtonWasPressed()
}
