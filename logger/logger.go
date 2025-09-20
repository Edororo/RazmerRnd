package logger

import (
	"context"
	"log"
	"time"

	"github.com/Edororo/RazmerRnd/internal/repository"
)

func LogNewEntries(ctx context.Context, repo *repository.Repository) {
	// Начинаем с уже загруженных данных (их не логируем)
	lastProducts := repo.LoadedProducts
	lastCart := repo.LoadedCart
	lastOrders := repo.LoadedOrders

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("[Logger] Завершение работы")
			return
		case <-ticker.C:
			// ✅ ВЫЗЫВАЕМ методы со скобками, чтобы получить слайсы
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
					log.Printf("[LOG] Новая позиция в корзине: %s x%d\n", c.ProductId, c.Quantity)
				}
				lastCart = len(cart)
			}

			if len(orders) > lastOrders {
				for _, o := range orders[lastOrders:] {
					log.Printf("[LOG] Новый заказ #%s от %s (товаров: %d)\n",
						o.ID, o.Customer, len(o.Items))
				}
				lastOrders = len(orders)
			}
		}
	}
}
