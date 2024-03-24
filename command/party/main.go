package main

import "fmt"

func main() {
	remoteControl := NewRemoteControl()

	livingRoomLight := NewLight("Living Room")
	stereo := NewStereo("Living Room")
	tv := NewTV("Living Room")
	hottub := NewHottub()

	livingRoomLightOnCommand := NewLightOnCommand(*livingRoomLight)
	livingRoomLightOffCommand := NewLightOffCommand(*livingRoomLight)
	stereoOnWithCD := NewStereoOnWithCDCommand(*stereo)
	stereoOff := NewStereoOffCommand(*stereo)
	tvOn := NewTVOnCommand(*tv)
	tvOff := NewTVOffCommand(*tv)
	hottubOn := NewHottubOnCommand(*hottub)
	hottubOff := NewHottubOffCommand(*hottub)

	partyOn := []Command{livingRoomLightOnCommand, stereoOnWithCD, tvOn, hottubOn}
	partyOff := []Command{livingRoomLightOffCommand, stereoOff, tvOff, hottubOff}

	partyOnMacro := NewMacroCommand(partyOn)
	partyOffMacro := NewMacroCommand(partyOff)

	remoteControl.SetCommand(0, partyOnMacro, partyOffMacro)

	fmt.Println(remoteControl)
	fmt.Println("\n--- Pushing Macro On ---")
	remoteControl.OnButtonWasPushed(0)
	fmt.Println("\n--- Pushing Macro Off ---")
	remoteControl.OffButtonWasPushed(0)
}
