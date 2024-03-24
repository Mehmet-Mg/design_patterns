package main

func main() {
	cook := NewCook()
	waitress := NewWaitress()
	customer := NewCustomer(waitress)
	customer.CreateOrder(NewBurgerAndFriesOrder(cook))
	customer.Hungry()
}
