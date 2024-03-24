package main

type MacroCommand struct {
	commands []Command
}

func NewMacroCommand(commands []Command) *MacroCommand {
	return &MacroCommand{
		commands: commands,
	}
}

func (m *MacroCommand) Execute() {
	for _, command := range m.commands {
		command.Execute()
	}
}

func (m *MacroCommand) Undo() {
	for _, command := range m.commands {
		command.Undo()
	}
}
