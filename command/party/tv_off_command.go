package main

type TVOffCommand struct {
	tv TV
}

func NewTVOffCommand(tv TV) *TVOffCommand {
	return &TVOffCommand{
		tv: tv,
	}
}

func (t *TVOffCommand) Execute() {
	t.tv.Off()
	t.tv.SetInputChannel()
}

func (t *TVOffCommand) Undo() {
	t.tv.Off()
}
