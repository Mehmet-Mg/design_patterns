package main

import "fmt"

type GarageDoor struct {
}

func NewGarageDoor() *GarageDoor {
	return &GarageDoor{}
}

func (g GarageDoor) Up() {
	fmt.Println("Garage Door is Up")
}

func (g GarageDoor) Down() {
	fmt.Println("Garage Door is Down")
}

func (g GarageDoor) Stop() {
	fmt.Println("Garage Door is Stopped")
}

func (g GarageDoor) LightOn() {
	fmt.Println("Garage light is on")
}

func (g GarageDoor) LightOff() {
	fmt.Println("Garage light is on")
}
