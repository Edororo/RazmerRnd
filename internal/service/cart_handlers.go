package server

import (
	"encoding/json"
	"net/http"

	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/gorilla/mux"
)

// --- GET /api/cart ---
func (s *Server) handleGetCartItems(w http.ResponseWriter, r *http.Request) {
	items := s.repo.GetCartItems()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// --- GET /api/cart/{id} ---
func (s *Server) handleGetCartItemByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, c := range s.repo.GetCartItems() {
		if c.ProductId == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	http.Error(w, "cart item not found", http.StatusNotFound)
}

// --- POST /api/cart ---
func (s *Server) handleAddCartItem(w http.ResponseWriter, r *http.Request) {
	var c model.CartItem
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if c.ProductId == "" || c.Quantity <= 0 {
		http.Error(w, "missing or invalid required fields", http.StatusBadRequest)
		return
	}

	s.repo.AddCartItem(c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

// --- PUT /api/cart/{id} ---
func (s *Server) handleUpdateCartItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updated model.CartItem
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	items := s.repo.GetCartItems()
	found := false
	for i, c := range items {
		if c.ProductId == id {
			if updated.Price != 0 {
				c.Price = updated.Price
			}
			if updated.Quantity != 0 {
				c.Quantity = updated.Quantity
			}
			items[i] = c
			s.repo.UpdateCart(items)
			found = true
			json.NewEncoder(w).Encode(c)
			break
		}
	}
	if !found {
		http.Error(w, "cart item not found", http.StatusNotFound)
	}
}

// --- DELETE /api/cart/{id} ---
func (s *Server) handleDeleteCartItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	items := s.repo.GetCartItems()
	for i, c := range items {
		if c.ProductId == id {
			items = append(items[:i], items[i+1:]...)
			s.repo.UpdateCart(items)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "cart item not found", http.StatusNotFound)
}
