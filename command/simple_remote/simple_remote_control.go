package main

type SimpleRemoteControl struct {
	slot Command
}

func NewSimpleRemoteControl() *SimpleRemoteControl {
	return &SimpleRemoteControl{}
}

func (r *SimpleRemoteControl) SetCommand(command Command) {
	r.slot = command
}

func (r SimpleRemoteControl) ButtonWasPushed() {
	r.slot.Execute()
}
