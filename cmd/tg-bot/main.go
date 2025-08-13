package main

import (
	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/Edororo/RazmerRnd/internal/repository"
	"github.com/Edororo/RazmerRnd/internal/service"
	"github.com/Edororo/RazmerRnd/logger"
)

func main() {
	repo := repository.NewRepository()
	ch := make(chan model.Entity, 10)

	svc := service.NewService(ch)

	// Горутина-производитель
	go svc.ProduceData()

	// Горутина-потребитель (репозиторий)
	go func() {
		for e := range ch {
			repo.AddEntity(e)
		}
	}()

	// Горутина-логгер
	go logger.LogNewEntries(repo)

	select {} // блокируем main
}
