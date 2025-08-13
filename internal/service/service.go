package service

import (
	"fmt"
	"github.com/Edororo/RazmerRnd/internal/model"
	"time"
)

type Service struct {
	Output chan model.Entity
}

func NewService(output chan model.Entity) *Service {
	return &Service{Output: output}
}

func (s *Service) ProduceData() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		// Product
		p := model.Product{
			ID:    fmt.Sprintf("p-%d", time.Now().UnixNano()),
			Name:  "T-Shirt",
			Price: 19.99,
		}
		s.Output <- p

		// CartItem
		c := model.CartItem{
			ProductId: p.ID,
			Price:     p.Price,
			Quantity:  2,
		}
		s.Output <- c

		// Order
		o := model.Order{
			ID:        fmt.Sprintf("o-%d", time.Now().UnixNano()),
			Customer:  "Джон Сноу",
			Items:     []model.CartItem{c},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		s.Output <- o
	}
}
