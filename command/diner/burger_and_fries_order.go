package main

type BurgerAndFriesOrder struct {
	cook Cook
}

func NewBurgerAndFriesOrder(cook *Cook) *BurgerAndFriesOrder {
	return &BurgerAndFriesOrder{
		cook: *cook,
	}
}

func (b *BurgerAndFriesOrder) OrderUp() {
	b.cook.MakeBurger()
	b.cook.MakeFries()
}
