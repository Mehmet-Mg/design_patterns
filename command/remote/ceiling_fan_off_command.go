package main

type CeilingFanOffCommand struct {
	ceilingFan CeilingFan
}

func NewCeilingFanOffCommand(ceilingFan CeilingFan) *CeilingFanOffCommand {
	return &CeilingFanOffCommand{
		ceilingFan: ceilingFan,
	}
}

func (c CeilingFanOffCommand) Execute() {
	c.ceilingFan.Off()
}
