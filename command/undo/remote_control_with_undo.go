package main

import (
	"fmt"
	"reflect"
)

type RemoteControl struct {
	onCommands  []Command
	offCommands []Command
	undoCommand Command
}

func NewRemoteControl() *RemoteControl {
	onCommands := make([]Command, 7)
	offCommands := make([]Command, 7)
	undoCommand := &NoCommand{}

	noCommand := &NoCommand{}

	for i := range onCommands {
		onCommands[i] = noCommand
		offCommands[i] = noCommand
	}

	return &RemoteControl{
		onCommands:  onCommands,
		offCommands: offCommands,
		undoCommand: undoCommand,
	}
}

func (r *RemoteControl) SetCommand(slot int, onCommand Command, offCommand Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

func (r RemoteControl) OnButtonWasPushed(slot int) {
	r.onCommands[slot].Execute()
}

func (r *RemoteControl) OffButtonWasPushed(slot int) {
	r.offCommands[slot].Execute()
}

func (r RemoteControl) UndoButtonWasPressed() {
	r.undoCommand.Undo()
}

func (r RemoteControl) String() string {
	stringResult := "\n------ Remote Control -------\n"
	for i := range r.onCommands {
		stringResult = stringResult + fmt.Sprintf("[slot %d] %v \t %v\n", i, getName(r.onCommands[i]), getName(r.offCommands[i]))
	}
	return stringResult
}

func getName(variable any) string {
	return reflect.TypeOf(variable).Elem().Name()
}
