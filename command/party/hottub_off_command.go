package main

type HottubOffCommand struct {
	hottub Hottub
}

func NewHottubOffCommand(hottub Hottub) *HottubOffCommand {
	return &HottubOffCommand{
		hottub: hottub,
	}
}

func (h HottubOffCommand) Execute() {
	h.hottub.Cool()
	h.hottub.Off()
}

func (h HottubOffCommand) Undo() {
	h.hottub.On()
}
