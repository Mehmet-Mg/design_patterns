package main

type TVOnCommand struct {
	tv TV
}

func NewTVOnCommand(tv TV) *TVOnCommand {
	return &TVOnCommand{
		tv: tv,
	}
}

func (t *TVOnCommand) Execute() {
	t.tv.On()
	t.tv.SetInputChannel()
}

func (t *TVOnCommand) Undo() {
	t.tv.Off()
}
