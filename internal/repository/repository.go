package repository

import (
	"fmt"
	"github.com/Edororo/RazmerRnd/internal/model"
	"sync"
)

type Repository struct {
	muProducts sync.Mutex
	muCart     sync.Mutex
	muOrders   sync.Mutex

	products []model.Product
	cart     []model.CartItem
	orders   []model.Order
}

func NewRepository() *Repository {
	return &Repository{}
}

// Добавление с блокировкой только нужного слайса
func (r *Repository) AddEntity(e model.Entity) {
	switch v := e.(type) {
	case model.Product:
		r.muProducts.Lock()
		r.products = append(r.products, v)
		r.muProducts.Unlock()
	case model.CartItem:
		r.muCart.Lock()
		r.cart = append(r.cart, v)
		r.muCart.Unlock()
	case model.Order:
		r.muOrders.Lock()
		r.orders = append(r.orders, v)
		r.muOrders.Unlock()
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}

// Методы для получения копий слайсов
func (r *Repository) GetProducts() []model.Product {
	r.muProducts.Lock()
	defer r.muProducts.Unlock()
	cp := make([]model.Product, len(r.products))
	copy(cp, r.products)
	return cp
}

func (r *Repository) GetCart() []model.CartItem {
	r.muCart.Lock()
	defer r.muCart.Unlock()
	cp := make([]model.CartItem, len(r.cart))
	copy(cp, r.cart)
	return cp
}

func (r *Repository) GetOrders() []model.Order {
	r.muOrders.Lock()
	defer r.muOrders.Unlock()
	cp := make([]model.Order, len(r.orders))
	copy(cp, r.orders)
	return cp
}
