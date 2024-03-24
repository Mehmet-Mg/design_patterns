package main

import "fmt"

const (
	OFF int = iota
	LOW
	MEDIUM
	HIGH
)

type CeilingFan struct {
	location string
	level    int
}

func NewCeilingFan(location string) *CeilingFan {
	return &CeilingFan{
		location: location,
	}
}

func (c *CeilingFan) High() {
	c.level = HIGH
	fmt.Println(c.location + " ceiling fan is on high")
}

func (c *CeilingFan) Medium() {
	c.level = MEDIUM
	fmt.Println(c.location + " ceiling fan is on medium")
}

func (c *CeilingFan) Low() {
	c.level = LOW
	fmt.Println(c.location + " ceiling fan is on low")
}

func (c *CeilingFan) Off() {
	c.level = OFF
	fmt.Println(c.location + " ceiling fan is on off")
}

func (c CeilingFan) GetSpeed() int {
	return c.level
}
