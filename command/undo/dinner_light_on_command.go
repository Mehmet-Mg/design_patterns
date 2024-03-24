package main

type DinnerLightOnCommand struct {
	light Light
}

func NewDinnerLightOnCommand(ligh Light) *DinnerLightOnCommand {
	return &DinnerLightOnCommand{
		light: ligh,
	}
}

func (l DinnerLightOnCommand) Execute() {
	l.light.On()
}

func (l DinnerLightOnCommand) Undo() {
	l.light.Off()
}
