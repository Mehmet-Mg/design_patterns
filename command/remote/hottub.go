package main

import "fmt"

type Hottub struct {
	on          bool
	temperature int
}

func NewHottub() *Hottub {
	return &Hottub{}
}

func (h *Hottub) On() {
	h.on = true
}

func (h *Hottub) Off() {
	h.on = false
}

func (h Hottub) BubblesOn() {
	if h.on {
		fmt.Println("Hottub is bubbling!")
	}
}

func (h Hottub) BubblesOff() {
	if h.on {
		fmt.Println("Hottub is not bubbling!")
	}
}

func (h Hottub) JetsOn() {
	if h.on {
		fmt.Println("Hottub jets are on.")
	}
}

func (h Hottub) JetsOff() {
	if h.on {
		fmt.Println("Hottub jets are off.")
	}
}

func (h *Hottub) SetTemperature(temperature int) {
	h.temperature = temperature
}

func (h *Hottub) Heat() {
	h.temperature = 105
	fmt.Println("Hottub is heating to a steaming 105 degrees")
}

func (h *Hottub) Cool() {
	h.temperature = 98
	fmt.Println("Hottub is cooling to 98 degrees")
}
