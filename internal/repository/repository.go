package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/Edororo/RazmerRnd/internal/model"
)

type Repository struct {
	muProducts sync.RWMutex
	muCart     sync.RWMutex
	muOrders   sync.RWMutex

	products []model.Product
	cart     []model.CartItem
	orders   []model.Order

	productsFile string
	cartFile     string
	ordersFile   string

	LoadedProducts int
	LoadedCart     int
	LoadedOrders   int
}

func NewRepository(productsFile, cartFile, ordersFile string) *Repository {
	r := &Repository{
		productsFile: productsFile,
		cartFile:     cartFile,
		ordersFile:   ordersFile,
	}

	// загрузка данных из файлов при старте
	_ = r.loadProducts()
	_ = r.loadCart()
	_ = r.loadOrders()

	r.LoadedProducts = len(r.products)
	r.LoadedCart = len(r.cart)
	r.LoadedOrders = len(r.orders)

	return r
}

func (r *Repository) AddProduct(p model.Product) {
	r.muProducts.Lock()
	defer r.muProducts.Unlock()
	r.products = append(r.products, p)
	_ = r.saveProducts()
}

func (r *Repository) GetProducts() []model.Product {
	r.muProducts.RLock()
	defer r.muProducts.RUnlock()
	cp := make([]model.Product, len(r.products))
	copy(cp, r.products)
	return cp
}

func (r *Repository) GetProductByID(id string) (*model.Product, error) {
	r.muProducts.RLock()
	defer r.muProducts.RUnlock()
	for _, p := range r.products {
		if p.ID == id {
			cp := p
			return &cp, nil
		}
	}
	return nil, errors.New("product not found")
}

func (r *Repository) UpdateProduct(id string, updated model.Product) error {
	r.muProducts.Lock()
	defer r.muProducts.Unlock()
	for i, p := range r.products {
		if p.ID == id {
			r.products[i] = updated
			return r.saveProducts()
		}
	}
	return errors.New("product not found")
}

func (r *Repository) DeleteProduct(id string) error {
	r.muProducts.Lock()
	defer r.muProducts.Unlock()
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return r.saveProducts()
		}
	}
	return errors.New("product not found")
}

func (r *Repository) AddCartItem(c model.CartItem) {
	r.muCart.Lock()
	defer r.muCart.Unlock()
	r.cart = append(r.cart, c)
	_ = r.saveCart()
}

func (r *Repository) GetCart() []model.CartItem {
	r.muCart.RLock()
	defer r.muCart.RUnlock()
	cp := make([]model.CartItem, len(r.cart))
	copy(cp, r.cart)
	return cp
}

func (r *Repository) GetCartItemByID(id string) (*model.CartItem, error) {
	r.muCart.RLock()
	defer r.muCart.RUnlock()
	for _, c := range r.cart {
		if c.ProductId == id {
			cp := c
			return &cp, nil
		}
	}
	return nil, errors.New("cart item not found")
}

func (r *Repository) UpdateCartItem(id string, updated model.CartItem) error {
	r.muCart.Lock()
	defer r.muCart.Unlock()
	for i, c := range r.cart {
		if c.ProductId == id {
			r.cart[i] = updated
			return r.saveCart()
		}
	}
	return errors.New("cart item not found")
}

func (r *Repository) DeleteCartItem(id string) error {
	r.muCart.Lock()
	defer r.muCart.Unlock()
	for i, c := range r.cart {
		if c.ProductId == id {
			r.cart = append(r.cart[:i], r.cart[i+1:]...)
			return r.saveCart()
		}
	}
	return errors.New("cart item not found")
}

func (r *Repository) AddOrder(o model.Order) {
	r.muOrders.Lock()
	defer r.muOrders.Unlock()
	r.orders = append(r.orders, o)
	_ = r.saveOrders()
}

func (r *Repository) GetOrders() []model.Order {
	r.muOrders.RLock()
	defer r.muOrders.RUnlock()
	cp := make([]model.Order, len(r.orders))
	copy(cp, r.orders)
	return cp
}

func (r *Repository) GetOrderByID(id string) (*model.Order, error) {
	r.muOrders.RLock()
	defer r.muOrders.RUnlock()
	for _, o := range r.orders {
		if o.ID == id {
			cp := o
			return &cp, nil
		}
	}
	return nil, errors.New("order not found")
}

func (r *Repository) UpdateOrder(id string, updated model.Order) error {
	r.muOrders.Lock()
	defer r.muOrders.Unlock()
	for i, o := range r.orders {
		if o.ID == id {
			r.orders[i] = updated
			return r.saveOrders()
		}
	}
	return errors.New("order not found")
}

func (r *Repository) DeleteOrder(id string) error {
	r.muOrders.Lock()
	defer r.muOrders.Unlock()
	for i, o := range r.orders {
		if o.ID == id {
			r.orders = append(r.orders[:i], r.orders[i+1:]...)
			return r.saveOrders()
		}
	}
	return errors.New("order not found")
}

func (r *Repository) saveProducts() error {
	return saveToFile(r.productsFile, r.products)
}

func (r *Repository) saveCart() error {
	return saveToFile(r.cartFile, r.cart)
}

func (r *Repository) saveOrders() error {
	return saveToFile(r.ordersFile, r.orders)
}

func (r *Repository) loadProducts() error {
	var data []model.Product
	if err := loadFromFile(r.productsFile, &data); err != nil {
		return err
	}
	r.muProducts.Lock()
	r.products = data
	r.muProducts.Unlock()
	return nil
}

func (r *Repository) loadCart() error {
	var data []model.CartItem
	if err := loadFromFile(r.cartFile, &data); err != nil {
		return err
	}
	r.muCart.Lock()
	r.cart = data
	r.muCart.Unlock()
	return nil
}

func (r *Repository) loadOrders() error {
	var data []model.Order
	if err := loadFromFile(r.ordersFile, &data); err != nil {
		return err
	}
	r.muOrders.Lock()
	r.orders = data
	r.muOrders.Unlock()
	return nil
}

func saveToFile(filename string, data interface{}) error {
	tmp := filename + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return os.Rename(tmp, filename)
}

func loadFromFile(filename string, target interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(target); err != nil {
		return fmt.Errorf("decode %s: %w", filename, err)
	}
	return nil
}
