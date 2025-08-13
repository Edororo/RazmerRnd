package model

import "time"

// Общий интерфейс
type Entity interface {
	GetID() string
}

type Order struct {
	ID         string     `json:"id"`
	Customer   string     `json:"customer"`
	Address    string     `json:"address"`
	Phone      string     `json:"phone"`
	Items      []CartItem `json:"items"`
	TotalPrice float64    `json:"total_price"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func (o Order) GetID() string { return o.ID }

type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Sizes       []string `json:"sizes"`
	ImageURL    string   `json:"image_url"`
}

func (p Product) GetID() string { return p.ID }

type CartItem struct {
	ProductId string  `json:"product_id"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

func (c CartItem) GetID() string { return c.ProductId }
