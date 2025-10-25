package handler

import (
	"encoding/json"
	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) HandleGetCartItems(w http.ResponseWriter, r *http.Request) {
	cart := h.service.GetCart()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (h *Handler) HandleGetCartItemByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	item, err := h.service.GetCartItemByID(id)
	if err != nil {
		http.Error(w, "cart item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) HandleAddCartItem(w http.ResponseWriter, r *http.Request) {
	var c model.CartItem
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if c.ProductId == "" || c.Quantity == 0 || c.Price == 0 {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}
	h.service.AddCartItem(c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (h *Handler) HandleUpdateCartItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updated model.CartItem
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateCartItem(id, updated); err != nil {
		http.Error(w, "cart item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) HandleDeleteCartItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeleteCartItem(id); err != nil {
		http.Error(w, "cart item not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
