package main

type LightOnCommand struct {
	light Light
}

func NewLightOnCommand(ligh Light) *LightOnCommand {
	return &LightOnCommand{
		light: ligh,
	}
}

func (l LightOnCommand) Execute() {
	l.light.On()
}

func (l LightOnCommand) Undo() {
	l.light.Off()
}
