package main

type LivingRoomLightOnCommand struct {
	light Light
}

func NewLivingRoomLightOnCommand(ligh Light) *LivingRoomLightOnCommand {
	return &LivingRoomLightOnCommand{
		light: ligh,
	}
}

func (l LivingRoomLightOnCommand) Execute() {
	l.light.On()
}

func (l LivingRoomLightOnCommand) Undo() {
	l.light.Off()
}
