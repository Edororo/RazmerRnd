package server

import (
	"github.com/Edororo/RazmerRnd/internal/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	repo *repository.Repository
	mux  *mux.Router
}

func NewServer(repo *repository.Repository) *Server {
	s := &Server{
		repo: repo,
		mux:  mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	// --- Products ---
	s.mux.HandleFunc("/api/products", s.handleGetProductByID).Methods("GET")
	s.mux.HandleFunc("/api/product/{id}", s.handleGetProductByID).Methods("GET")
	s.mux.HandleFunc("/api/product", s.handleAddProduct).Methods("POST")
	s.mux.HandleFunc("/api/product/{id}", s.handleUpdateProduct).Methods("PUT")
	s.mux.HandleFunc("/api/product/{id}", s.handleDeleteProduct).Methods("DELETE")

	// --- Cart ---
	s.mux.HandleFunc("/api/cart", s.handleGetCartItems).Methods("GET")
	s.mux.HandleFunc("/api/cart/{id}", s.handleGetCartItemByID).Methods("GET")
	s.mux.HandleFunc("/api/cart", s.handleAddCartItem).Methods("POST")
	s.mux.HandleFunc("/api/cart/{id}", s.handleUpdateCartItem).Methods("PUT")
	s.mux.HandleFunc("/api/cart/{id}", s.handleDeleteCartItem).Methods("DELETE")

	// --- Orders ---
	s.mux.HandleFunc("/api/orders", s.handleGetOrders).Methods("GET")
	s.mux.HandleFunc("/api/order/{id}", s.handleGetOrderByID).Methods("GET")
	s.mux.HandleFunc("/api/order", s.handleAddOrder).Methods("POST")
	s.mux.HandleFunc("/api/order/{id}", s.handleUpdateOrder).Methods("PUT")
	s.mux.HandleFunc("/api/order/{id}", s.handleDeleteOrder).Methods("DELETE")
}

func (s *Server) Run(addr string) {
	log.Printf("Webserver запущен на %s", addr)
	log.Fatal(http.ListenAndServe(addr, s.mux))
}
