package main

type Waitress struct {
	order Order
}

func NewWaitress() *Waitress {
	return &Waitress{}
}

func (w *Waitress) TakeOrder(order Order) {
	w.order = order
	w.order.OrderUp()
}
