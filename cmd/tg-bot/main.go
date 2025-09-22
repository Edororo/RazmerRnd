package main

import (
	"context"
	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/Edororo/RazmerRnd/internal/repository"
	"github.com/Edororo/RazmerRnd/internal/service"
	"github.com/Edororo/RazmerRnd/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	repo := repository.NewRepository(
		"data/products.json",
		"data/cart.json",
		"data/orders.json",
	)

	ch := make(chan model.Entity, 10)

	svc := service.NewService(ch)

	// 5️⃣ Запуск горутин

	go svc.ProduceData(ctx)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("[Consumer] Завершение работы")
				close(ch)
				return
			case e := <-ch:
				repo.AddEntity(e)
			}
		}
	}()

	go logger.LogNewEntries(ctx, repo)

	sig := <-sigCh
	log.Printf("Получен сигнал %s. Завершаем работу...", sig)
	cancel()

	time.Sleep(500 * time.Millisecond)
	log.Println("Приложение завершено.")
}
