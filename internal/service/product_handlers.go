package server

import (
	"encoding/json"
	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.repo.GetProducts())
}

func (s *Server) handleGetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, p := range s.repo.GetProducts() {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "product not found", http.StatusNotFound)
}

func (s *Server) handleAddProduct(w http.ResponseWriter, r *http.Request) {
	var p model.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if p.ID == "" || p.Name == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}
	s.repo.AddProduct(p)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func (s *Server) handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updated model.Product
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	products := s.repo.GetProducts()
	for i, p := range products {
		if p.ID == id {
			if updated.Name != "" {
				p.Name = updated.Name
			}
			if updated.Price != 0 {
				p.Price = updated.Price
			}
			products[i] = p
			s.repo.UpdateProducts(products)
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "product not found", http.StatusNotFound)
}

func (s *Server) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	products := s.repo.GetProducts()
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			s.repo.UpdateProducts(products)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "product not found", http.StatusNotFound)
}
