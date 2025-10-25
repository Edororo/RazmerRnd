package handler

import (
	"encoding/json"
	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) HandleGetOrders(w http.ResponseWriter, r *http.Request) {
	orders := h.service.GetOrders()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *Handler) HandleGetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	order, err := h.service.GetOrderByID(id)
	if err != nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order)
}

func (h *Handler) HandleAddOrder(w http.ResponseWriter, r *http.Request) {
	var o model.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if o.ID == "" || o.Customer == "" || len(o.Items) == 0 {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}
	h.service.AddOrder(o)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(o)
}

func (h *Handler) HandleUpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updated model.Order
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateOrder(id, updated); err != nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) HandleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeleteOrder(id); err != nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
