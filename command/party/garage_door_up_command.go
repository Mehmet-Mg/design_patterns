package main

type GarageDoorUpCommand struct {
	garageDoor GarageDoor
}

func NewGarageDoorUpCommand(garageDoor GarageDoor) *GarageDoorUpCommand {
	return &GarageDoorUpCommand{
		garageDoor: garageDoor,
	}
}

func (g GarageDoorUpCommand) Execute() {
	g.garageDoor.Up()
}

func (g GarageDoorUpCommand) Undo() {
	g.garageDoor.Down()
}
