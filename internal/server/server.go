package server

import (
	"github.com/Edororo/RazmerRnd/internal/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	handler *handler.Handler
	mux     *mux.Router
}

func NewServer(h *handler.Handler) *Server {
	s := &Server{
		handler: h,
		mux:     mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/api/products", s.handler.HandleGetProducts).Methods("GET")
	s.mux.HandleFunc("/api/product/{id}", s.handler.HandleGetProductByID).Methods("GET")
	s.mux.HandleFunc("/api/product", s.handler.HandleAddProduct).Methods("POST")
	s.mux.HandleFunc("/api/product/{id}", s.handler.HandleUpdateProduct).Methods("PUT")
	s.mux.HandleFunc("/api/product/{id}", s.handler.HandleDeleteProduct).Methods("DELETE")

	s.mux.HandleFunc("/api/cart", s.handler.HandleGetCartItems).Methods("GET")
	s.mux.HandleFunc("/api/cart/{id}", s.handler.HandleGetCartItemByID).Methods("GET")
	s.mux.HandleFunc("/api/cart", s.handler.HandleAddCartItem).Methods("POST")
	s.mux.HandleFunc("/api/cart/{id}", s.handler.HandleUpdateCartItem).Methods("PUT")
	s.mux.HandleFunc("/api/cart/{id}", s.handler.HandleDeleteCartItem).Methods("DELETE")

	s.mux.HandleFunc("/api/orders", s.handler.HandleGetOrders).Methods("GET")
	s.mux.HandleFunc("/api/order/{id}", s.handler.HandleGetOrderByID).Methods("GET")
	s.mux.HandleFunc("/api/order", s.handler.HandleAddOrder).Methods("POST")
	s.mux.HandleFunc("/api/order/{id}", s.handler.HandleUpdateOrder).Methods("PUT")
	s.mux.HandleFunc("/api/order/{id}", s.handler.HandleDeleteOrder).Methods("DELETE")
}

func (s *Server) Run(addr string) {
	log.Printf("Server running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, s.mux))
}
