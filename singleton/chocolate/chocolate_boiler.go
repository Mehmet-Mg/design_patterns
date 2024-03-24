package main

import "fmt"

type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

var instance *ChocolateBoiler

func GetInstance() *ChocolateBoiler {
	if instance == nil {
		fmt.Println("Creating unique instance of Chocolate Boiler")
		instance = &ChocolateBoiler{
			empty:  true,
			boiled: false,
		}
	}
	fmt.Println("Returning instance of Chocolate Boiler")

	return instance
}

func (c *ChocolateBoiler) Fill() {
	if c.empty {
		c.empty = false
		c.boiled = false
	}
}

func (c *ChocolateBoiler) Drain() {
	if !c.empty && c.boiled {
		c.empty = true
	}
}

func (c *ChocolateBoiler) Boil() {
	if !c.empty && !c.boiled {
		c.boiled = true
	}
}
