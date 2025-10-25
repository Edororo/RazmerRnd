package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Edororo/RazmerRnd/internal/model"
	"github.com/Edororo/RazmerRnd/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Products

// @Summary Get all products
// @Tags products
// @Produce json
// @Success 200 {array} model.Product
// @Router /api/products [get]
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := h.service.GetProducts()
	writeJSON(w, products)
}

// @Summary Get product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} model.Product
// @Failure 404 {string} string "product not found"
// @Router /api/products/{id} [get]
func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	p, err := h.service.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, p)
}

// @Summary Add product
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product"
// @Success 201 {object} model.Product
// @Failure 400 {string} string "bad request"
// @Router /api/product [post]
func (h *Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var p model.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if p.ID == "" || p.Name == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}
	h.service.AddProduct(p)
	writeJSON(w, p)
}

// @Summary Update product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body model.Product true "Product"
// @Success 200 {object} model.Product
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "product not found"
// @Router /api/product/{id} [put]
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p model.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateProduct(id, p); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, p)
}

// @Summary Delete product
// @Tags products
// @Param id path string true "Product ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {string} string "product not found"
// @Router /api/product/{id} [delete]
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeleteProduct(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Cart

// @Summary Get all cart items
// @Tags cart
// @Produce json
// @Success 200 {array} model.CartItem
// @Router /api/cart [get]
func (h *Handler) GetCartItems(w http.ResponseWriter, r *http.Request) {
	items := h.service.GetCart()
	writeJSON(w, items)
}

// @Summary Get cart item by ID
// @Tags cart
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} model.CartItem
// @Failure 404 {string} string "cart item not found"
// @Router /api/cart/{id} [get]
func (h *Handler) GetCartItemByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	c, err := h.service.GetCartItemByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, c)
}

// @Summary Add cart item
// @Tags cart
// @Accept json
// @Produce json
// @Param cartItem body model.CartItem true "CartItem"
// @Success 201 {object} model.CartItem
// @Failure 400 {string} string "bad request"
// @Router /api/cart [post]
func (h *Handler) AddCartItem(w http.ResponseWriter, r *http.Request) {
	var c model.CartItem
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if c.ProductId == "" {
		http.Error(w, "missing product_id", http.StatusBadRequest)
		return
	}
	h.service.AddCartItem(c)
	writeJSON(w, c)
}

// @Summary Update cart item
// @Tags cart
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param cartItem body model.CartItem true "CartItem"
// @Success 200 {object} model.CartItem
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "cart item not found"
// @Router /api/cart/{id} [put]
func (h *Handler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var c model.CartItem
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateCartItem(id, c); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, c)
}

// @Summary Delete cart item
// @Tags cart
// @Param id path string true "Product ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {string} string "cart item not found"
// @Router /api/cart/{id} [delete]
func (h *Handler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeleteCartItem(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Orders

// @Summary Get all orders
// @Tags orders
// @Produce json
// @Success 200 {array} model.Order
// @Router /api/orders [get]
func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders := h.service.GetOrders()
	writeJSON(w, orders)
}

// @Summary Get order by ID
// @Tags orders
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} model.Order
// @Failure 404 {string} string "order not found"
// @Router /api/orders/{id} [get]
func (h *Handler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	o, err := h.service.GetOrderByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, o)
}

// @Summary Add order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body model.Order true "Order"
// @Success 201 {object} model.Order
// @Failure 400 {string} string "bad request"
// @Router /api/order [post]
func (h *Handler) AddOrder(w http.ResponseWriter, r *http.Request) {
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
	writeJSON(w, o)
}

// @Summary Update order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body model.Order true "Order"
// @Success 200 {object} model.Order
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "order not found"
// @Router /api/order/{id} [put]
func (h *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var o model.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateOrder(id, o); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, o)
}

// @Summary Delete order
// @Tags orders
// @Param id path string true "Order ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {string} string "order not found"
// @Router /api/order/{id} [delete]
func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeleteOrder(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
