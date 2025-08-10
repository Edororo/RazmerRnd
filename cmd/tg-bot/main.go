package main

import (
	"fmt"
	"github.com/Edororo/RazmerRnd/internal/repository"
	"github.com/Edororo/RazmerRnd/internal/service"
)

func main() {
	fmt.Println("Начало генерации данных")

	service.GenerateDataOnce()
	service.GenerateDataOnce()

	fmt.Println("Все данные сохранены")

	// Пример вывода заказов
	fmt.Println("Заказы")
	for _, order := range repository.GetOrders() {
		fmt.Printf("%+v\n", order)
	}
}
