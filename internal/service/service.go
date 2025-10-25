package service

import (
	"time"

	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/Edororo/RazmerRnd/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetProducts() []model.Product {
	return s.repo.GetProducts()
}

func (s *Service) GetProductByID(id string) (*model.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *Service) AddProduct(p model.Product) {
	s.repo.AddProduct(p)
}

func (s *Service) UpdateProduct(id string, updated model.Product) error {
	return s.repo.UpdateProduct(id, updated)
}

func (s *Service) DeleteProduct(id string) error {
	return s.repo.DeleteProduct(id)
}

func (s *Service) GetCart() []model.CartItem {
	return s.repo.GetCart()
}

func (s *Service) GetCartItemByID(id string) (*model.CartItem, error) {
	return s.repo.GetCartItemByID(id)
}

func (s *Service) AddCartItem(c model.CartItem) {
	s.repo.AddCartItem(c)
}

func (s *Service) UpdateCartItem(id string, updated model.CartItem) error {
	return s.repo.UpdateCartItem(id, updated)
}

func (s *Service) DeleteCartItem(id string) error {
	return s.repo.DeleteCartItem(id)
}

func (s *Service) GetOrders() []model.Order {
	return s.repo.GetOrders()
}

func (s *Service) GetOrderByID(id string) (*model.Order, error) {
	return s.repo.GetOrderByID(id)
}

func (s *Service) AddOrder(o model.Order) {
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	s.repo.AddOrder(o)
}

func (s *Service) UpdateOrder(id string, updated model.Order) error {
	updated.UpdatedAt = time.Now()
	return s.repo.UpdateOrder(id, updated)
}

func (s *Service) DeleteOrder(id string) error {
	return s.repo.DeleteOrder(id)
}
