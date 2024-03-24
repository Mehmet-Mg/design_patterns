package main

type LightOffCommand struct {
	light Light
}

func NewLightOffCommand(ligh Light) *LightOffCommand {
	return &LightOffCommand{
		light: ligh,
	}
}

func (l LightOffCommand) Execute() {
	l.light.Off()
}

func (l LightOffCommand) Undo() {
	l.light.On()
}
