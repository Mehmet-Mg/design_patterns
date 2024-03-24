package main

type Customer struct {
	waitress Waitress
	order    Order
}

func NewCustomer(waitress *Waitress) *Customer {
	return &Customer{
		waitress: *waitress,
	}
}

func (c *Customer) CreateOrder(order Order) {
	c.order = order
}

func (c *Customer) Hungry() {
	c.waitress.TakeOrder(c.order)
}
