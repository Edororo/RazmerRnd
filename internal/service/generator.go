package service

import (
	"fmt"
	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/Edororo/RazmerRnd/internal/repository"
	"time"
)

// в будущем буду заменять на ввод через тг бота

func GenerateDataOnce() {
	// Создание CartItem
	item := model.CartItem{
		ProductId: fmt.Sprintf("product-%d", time.Now().UnixNano()),
		Price:     199.99,
		Quantity:  2,
	}
	repository.SaveToStackable(item)

	// Создание Order
	order := model.Order{
		ID:         fmt.Sprintf("Заказ-%d", time.Now().UnixNano()),
		Phone:      "+70000000000",
		Address:    "Город, Улица, Дом",
		Items:      []model.CartItem{item},
		TotalPrice: item.Price * float64(item.Quantity),
		Status:     "Created",
		Client:     "Имя Клиента",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	// передача в репозиторий
	repository.SaveToStackable(order)
}
