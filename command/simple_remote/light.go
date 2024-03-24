package main

import "fmt"

type Light struct {
}

func NewLight() *Light {
	return &Light{}
}

func (l Light) On() {
	fmt.Println("The light is on")
}

func (l Light) Off() {
	fmt.Println("The light is off")
}
