package service

import (
	"context"
	"time"

	"github.com/Edororo/RazmerRnd/internal/model"
)

type Service struct {
	ch chan<- model.Entity
}

func NewService(ch chan<- model.Entity) *Service {
	return &Service{ch: ch}
}

// Производит новые данные каждые 2 секунды
func (s *Service) ProduceData(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Пример тестовых данных
			p := model.Product{ID: "p1", Name: "Sneakers", Price: 9990}
			s.ch <- p

			c := model.CartItem{ProductId: "p1", Price: 9990, Quantity: 1}
			s.ch <- c

			o := model.Order{ID: "o1", Customer: "John Doe", Items: []model.CartItem{c}}
			s.ch <- o
		}
	}
}
