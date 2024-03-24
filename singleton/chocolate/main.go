package main

func main() {
	boiler := GetInstance()
	boiler.Fill()
	boiler.Boil()
	boiler.Drain()

	boiler2 := GetInstance()
	boiler2.Fill()
}
