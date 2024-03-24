package main

import "fmt"

type GarageDoor struct {
	location string
}

func NewGarageDoor(location string) *GarageDoor {
	return &GarageDoor{
		location: location,
	}
}

func (g GarageDoor) Up() {
	fmt.Println(g.location + " garage Door is Up")
}

func (g GarageDoor) Down() {
	fmt.Println(g.location + " garage Door is Down")
}

func (g GarageDoor) Stop() {
	fmt.Println(g.location + " garage Door is Stopped")
}

func (g GarageDoor) LightOn() {
	fmt.Println(g.location + " garage light is on")
}

func (g GarageDoor) LightOff() {
	fmt.Println(g.location + " garage light is on")
}
