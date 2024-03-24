package main

type CeilingFanOffCommand struct {
	ceilingFan CeilingFan
	prevSpeed  int
}

func NewCeilingFanOffCommand(ceilingFan CeilingFan) *CeilingFanOffCommand {
	return &CeilingFanOffCommand{
		ceilingFan: ceilingFan,
	}
}

func (c CeilingFanOffCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.Off()
}

func (c CeilingFanOffCommand) Undo() {
	switch c.prevSpeed {
	case HIGH:
		c.ceilingFan.High()
	case MEDIUM:
		c.ceilingFan.Medium()
	case LOW:
		c.ceilingFan.Low()
	default:
		c.ceilingFan.Off()
	}
}
