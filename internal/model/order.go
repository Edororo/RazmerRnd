package model

import "time"

type CartItem struct {
	ProductId string
	Price     float64
	Quantity  int
}

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
