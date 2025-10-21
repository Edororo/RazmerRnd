package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/Edororo/RazmerRnd/internal/repository"
)

func main() {
	// Создаём репозиторий
	repo := repository.NewRepository(
		"data/products.json",
		"data/cart.json",
		"data/orders.json",
	)

	// Контекст для graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Слушаем сигналы ОС для остановки
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	// --- Web server ---
	mux := http.NewServeMux()

	// --- Products ---
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, repo.GetProducts())
		case http.MethodPost:
			var p model.Product
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if p.ID == "" || p.Name == "" {
				http.Error(w, "missing required fields", http.StatusBadRequest)
				return
			}
			repo.AddProduct(p)
			writeJSON(w, p)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/api/products/"):]
		switch r.Method {
		case http.MethodGet:
			p, err := repo.GetProductByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			writeJSON(w, p)
		case http.MethodPut:
			var p model.Product
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := repo.UpdateProduct(id, p); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			writeJSON(w, p)
		case http.MethodDelete:
			if err := repo.DeleteProduct(id); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// --- Cart ---
	mux.HandleFunc("/api/cart", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, repo.GetCart())
		case http.MethodPost:
			var c model.CartItem
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if c.ProductId == "" {
				http.Error(w, "missing product_id", http.StatusBadRequest)
				return
			}
			repo.AddCartItem(c)
			writeJSON(w, c)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/cart/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/api/cart/"):]
		switch r.Method {
		case http.MethodGet:
			c, err := repo.GetCartItemByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			writeJSON(w, c)
		case http.MethodPut:
			var c model.CartItem
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := repo.UpdateCartItem(id, c); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			writeJSON(w, c)
		case http.MethodDelete:
			if err := repo.DeleteCartItem(id); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// --- Orders ---
	mux.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, repo.GetOrders())
		case http.MethodPost:
			var o model.Order
			if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if o.ID == "" || o.Customer == "" {
				http.Error(w, "missing required fields", http.StatusBadRequest)
				return
			}
			repo.AddOrder(o)
			writeJSON(w, o)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/orders/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/api/orders/"):]
		switch r.Method {
		case http.MethodGet:
			o, err := repo.GetOrderByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			writeJSON(w, o)
		case http.MethodPut:
			var o model.Order
			if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := repo.UpdateOrder(id, o); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			writeJSON(w, o)
		case http.MethodDelete:
			if err := repo.DeleteOrder(id); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// --- HTTP Server ---
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Запуск сервера в отдельной горутине
	go func() {
		log.Println("Server started on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	// Ждём сигнал остановки
	<-stop
	log.Println("Shutdown signal received")

	// Graceful shutdown
	ctxShutdown, cancelShutdown := context.WithTimeout(ctx, 5*time.Second)
	defer cancelShutdown()
	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Server exited properly")
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
