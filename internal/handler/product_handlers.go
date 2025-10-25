package handler

import (
	"encoding/json"
	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	products := h.service.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *Handler) HandleGetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	product, err := h.service.GetProductByID(id)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *Handler) HandleAddProduct(w http.ResponseWriter, r *http.Request) {
	var p model.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if p.ID == "" || p.Name == "" || p.Price == 0 {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}
	h.service.AddProduct(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func (h *Handler) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updated model.Product
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateProduct(id, updated); err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeleteProduct(id); err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
