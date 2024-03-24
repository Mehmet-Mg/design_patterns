package main

type CeilingFanHighCommand struct {
	ceilingFan CeilingFan
	prevSpeed  int
}

func NewCeilingFanHighCommand(ceilingFan CeilingFan) *CeilingFanHighCommand {
	return &CeilingFanHighCommand{
		ceilingFan: ceilingFan,
	}
}

func (c CeilingFanHighCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.High()
}

func (c CeilingFanHighCommand) Undo() {
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
