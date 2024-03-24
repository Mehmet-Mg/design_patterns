package main

func main() {
	remote := NewSimpleRemoteControl()
	light := NewLight()
	garageDoor := NewGarageDoor()

	lightOn := NewLightOnCommand(*light)
	garageOpen := NewGarageDoorOpenCommand(*garageDoor)

	remote.SetCommand(lightOn)
	remote.ButtonWasPushed()
	remote.SetCommand(garageOpen)
	remote.ButtonWasPushed()
}
