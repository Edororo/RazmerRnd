package logger

import (
	"context"
	"github.com/Edororo/RazmerRnd/internal/repository"
	"log"
	"time"
)

func LogNewEntries(ctx context.Context, repo *repository.Repository) {
	var lastProducts, lastCart, lastOrders int

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("[Logger] Завершение работы")
			return
		case <-ticker.C:
			products := repo.GetProducts()
			cart := repo.GetCart()
			orders := repo.GetOrders()

			if len(products) > lastProducts {
				for _, p := range products[lastProducts:] {
					log.Printf("[LOG] Новый товар: %s (%.2f ₽)\n", p.Name, p.Price)
				}
				lastProducts = len(products)
			}

			if len(cart) > lastCart {
				for _, c := range cart[lastCart:] {
					log.Printf("[LOG] В корзину добавлено: %s x%d\n", c.ProductId, c.Quantity)
				}
				lastCart = len(cart)
			}

			if len(orders) > lastOrders {
				for _, o := range orders[lastOrders:] {
					log.Printf("[LOG] Новый заказ #%s от %s (товаров: %d)\n", o.ID, o.Customer, len(o.Items))
				}
				lastOrders = len(orders)
			}
		}
	}
}
