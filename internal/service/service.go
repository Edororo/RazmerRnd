package service

import (
	"context"
	"github.com/Edororo/RazmerRnd/internal/model"
	"time"
)

type Service struct {
	ch chan<- model.Entity
}

func NewService(ch chan<- model.Entity) *Service {
	return &Service{ch: ch}
}

func (s *Service) ProduceData(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Данные для примера
			p := model.Product{ID: "p1", Name: "T-Shirt", Price: 1999.00}
			s.ch <- p

			c := model.CartItem{ProductId: "p1", Price: 1999.00, Quantity: 1}
			s.ch <- c

			o := model.Order{ID: "o1", Customer: "John Doe", Items: []model.CartItem{c}}
			s.ch <- o
		}
	}
}
