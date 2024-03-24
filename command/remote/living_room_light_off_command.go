package main

type LivingRoomLightOffCommand struct {
	light Light
}

func NewLivingRoomLightOffCommand(ligh Light) *LivingRoomLightOffCommand {
	return &LivingRoomLightOffCommand{
		light: ligh,
	}
}

func (l LivingRoomLightOffCommand) Execute() {
	l.light.Off()
}
