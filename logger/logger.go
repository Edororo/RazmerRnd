package logger

import (
	"github.com/Edororo/RazmerRnd/internal/repository"
	"log"
	"time"
)

func LogNewEntries(repo *repository.Repository) {
	var lastProducts, lastCart, lastOrders int

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		products := repo.GetProducts()
		cart := repo.GetCart()
		orders := repo.GetOrders()

		if len(products) > lastProducts {
			log.Println("Новый продукт:", products[lastProducts:])
			lastProducts = len(products)
		}
		if len(cart) > lastCart {
			log.Println("Новые товары в корзине:", cart[lastCart:])
			lastCart = len(cart)
		}
		if len(orders) > lastOrders {
			log.Println("Новый заказ:", orders[lastOrders:])
			lastOrders = len(orders)
		}
	}
}
