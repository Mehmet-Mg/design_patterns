package main

import "fmt"

type Cook struct {
}

func NewCook() *Cook {
	return &Cook{}
}

func (c Cook) MakeBurger() {
	fmt.Println("Making a burger")
}

func (c Cook) MakeFries() {
	fmt.Println("Making fries")
}
