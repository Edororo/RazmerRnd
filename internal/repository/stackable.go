package repository

import (
	"fmt"
	"github.com/Edororo/RazmerRnd/internal/model"
)

var cartItems []model.CartItem
var orders []model.Order

// сохранение структуры в нужный слайс, принятие интерфейса
func SaveToStackable(s model.Stackable) {
	switch v := s.(type) {
	case model.Order:
		orders = append(orders, v)
		fmt.Println("Сохранённый заказ: ", v)
	case model.CartItem:
		cartItems = append(cartItems, v)
		fmt.Println("Товары в корзине: ", v)
	}
}

func GetCartItems() []model.CartItem {
	return cartItems
}

func GetOrders() []model.Order {
	return orders
}
