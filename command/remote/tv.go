package main

import "fmt"

type TV struct {
	location string
	channel  int
}

func NewTV(location string) *TV {
	return &TV{
		location: location,
	}
}

func (t *TV) On() {
	fmt.Println("TV is on")
}

func (t *TV) Off() {
	fmt.Println("TV is off")
}

func (t *TV) SetInputChannel() {
	t.channel = 3
	fmt.Println("Channel is set for VCR")
}
