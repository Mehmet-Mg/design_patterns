package main

import "fmt"

func main() {
	remoteControl := NewRemoteControl()

	livingRoomLight := NewLight("Living Room")
	kitchenLight := NewLight("Kitchen")
	ceilingFan := NewCeilingFan("Living Room")
	garageDoor := NewGarageDoor("Garage")
	stereo := NewStereo("Living Room")

	livingRoomLightOnCommand := NewLightOnCommand(*livingRoomLight)
	livingRoomLightOffCommand := NewLightOffCommand(*livingRoomLight)
	kitchenLightOnCommand := NewLightOnCommand(*kitchenLight)
	kitchenLightOffCommand := NewLightOffCommand(*kitchenLight)

	ceilingFanOn := NewCeilingFanOnCommand(*ceilingFan)
	ceilingFanOff := NewCeilingFanOffCommand(*ceilingFan)

	garageDoorUp := NewGarageDoorUpCommand(*garageDoor)
	garageDoorDown := NewGarageDoorDownCommand(*garageDoor)

	stereoOnWithCD := NewStereoOnWithCDCommand(*stereo)
	stereoOff := NewStereoOffCommand(*stereo)

	remoteControl.SetCommand(0, livingRoomLightOnCommand, livingRoomLightOffCommand)
	remoteControl.SetCommand(1, kitchenLightOnCommand, kitchenLightOffCommand)
	remoteControl.SetCommand(2, ceilingFanOn, ceilingFanOff)
	remoteControl.SetCommand(3, stereoOnWithCD, stereoOff)
	remoteControl.SetCommand(4, garageDoorUp, garageDoorDown)

	fmt.Println(remoteControl)

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	remoteControl.OnButtonWasPushed(1)
	remoteControl.OffButtonWasPushed(1)
	remoteControl.OnButtonWasPushed(2)
	remoteControl.OffButtonWasPushed(2)
	remoteControl.OnButtonWasPushed(3)
	remoteControl.OffButtonWasPushed(3)
	remoteControl.OnButtonWasPushed(4)
	remoteControl.OffButtonWasPushed(4)
}
