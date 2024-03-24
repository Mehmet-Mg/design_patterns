package main

type HottubOnCommand struct {
	hottub Hottub
}

func NewHottubOnCommand(hottub Hottub) *HottubOnCommand {
	return &HottubOnCommand{
		hottub: hottub,
	}
}

func (h HottubOnCommand) Execute() {
	h.hottub.On()
	h.hottub.Heat()
	h.hottub.BubblesOn()
}
