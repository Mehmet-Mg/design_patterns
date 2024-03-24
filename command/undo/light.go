package main

import "fmt"

type Light struct {
	location string
}

func NewLight(location string) *Light {
	return &Light{
		location: location,
	}
}

func (l Light) On() {
	fmt.Println(l.location + " light is on")
}

func (l Light) Off() {
	fmt.Println(l.location + " light is off")
}
