package model

import "time"

type Stackable interface {
	isStackable()
}
type CartItem struct {
	ProductId string
	Price     float64
	Quantity  int
}

func (c CartItem) isStackable() {}

type Order struct {
	ID         string
	Phone      string
	Address    string
	Items      []CartItem
	TotalPrice float64
	Status     string
	Client     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (o Order) isStackable() {}
