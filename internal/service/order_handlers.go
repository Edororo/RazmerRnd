package server

import (
	"encoding/json"
	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// --- GET /api/orders ---
func (s *Server) handleGetOrders(w http.ResponseWriter, r *http.Request) {
	orders := s.repo.GetOrders()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// --- GET /api/order/{id} ---
func (s *Server) handleGetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, o := range s.repo.GetOrders() {
		if o.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(o)
			return
		}
	}
	http.Error(w, "order not found", http.StatusNotFound)
}

// --- POST /api/order ---
func (s *Server) handleAddOrder(w http.ResponseWriter, r *http.Request) {
	var o model.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Проверяем обязательные поля
	if o.ID == "" || o.Customer == "" || len(o.Items) == 0 {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()

	s.repo.AddOrder(o)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(o)
}

// --- PUT /api/order/{id} ---
func (s *Server) handleUpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updated model.Order
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orders := s.repo.GetOrders()
	found := false
	for i, o := range orders {
		if o.ID == id {
			if updated.Customer != "" {
				o.Customer = updated.Customer
			}
			if updated.Status != "" {
				o.Status = updated.Status
			}
			if len(updated.Items) > 0 {
				o.Items = updated.Items
			}
			o.UpdatedAt = time.Now()
			orders[i] = o
			s.repo.UpdateOrders(orders)
			found = true
			json.NewEncoder(w).Encode(o)
			break
		}
	}
	if !found {
		http.Error(w, "order not found", http.StatusNotFound)
	}
}

// --- DELETE /api/order/{id} ---
func (s *Server) handleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	orders := s.repo.GetOrders()
	for i, o := range orders {
		if o.ID == id {
			orders = append(orders[:i], orders[i+1:]...)
			s.repo.UpdateOrders(orders)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "order not found", http.StatusNotFound)
}
