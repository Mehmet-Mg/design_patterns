package main

type Command interface {
	Execute()
	Undo()
}
