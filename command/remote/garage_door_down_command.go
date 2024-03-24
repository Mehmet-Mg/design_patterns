package main

type GarageDoorDownCommand struct {
	garageDoor GarageDoor
}

func NewGarageDoorDownCommand(garageDoor GarageDoor) *GarageDoorDownCommand {
	return &GarageDoorDownCommand{
		garageDoor: garageDoor,
	}
}

func (g GarageDoorDownCommand) Execute() {
	g.garageDoor.Down()
}
