package main

type CeilingFanOnCommand struct {
	ceilingFan CeilingFan
}

func NewCeilingFanOnCommand(ceilingFan CeilingFan) *CeilingFanOnCommand {
	return &CeilingFanOnCommand{
		ceilingFan: ceilingFan,
	}
}

func (c CeilingFanOnCommand) Execute() {
	c.ceilingFan.High()
}
