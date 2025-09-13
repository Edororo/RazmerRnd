package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/Edororo/RazmerRnd/internal/model"
)

type Repository struct {
	muProducts sync.Mutex
	muCart     sync.Mutex
	muOrders   sync.Mutex

	products []model.Product
	cart     []model.CartItem
	orders   []model.Order

	productsFile string
	cartFile     string
	ordersFile   string
}

// Создание репозитория с указанием файлов для хранения данных
func NewRepository(productsFile, cartFile, ordersFile string) *Repository {
	r := &Repository{
		productsFile: productsFile,
		cartFile:     cartFile,
		ordersFile:   ordersFile,
	}

	// Загружаем данные при старте
	r.loadProducts()
	r.loadCart()
	r.loadOrders()

	return r
}

// Универсальный метод
func (r *Repository) AddEntity(e model.Entity) {
	switch v := e.(type) {
	case model.Product:
		r.AddProduct(v)
	case model.CartItem:
		r.AddCartItem(v)
	case model.Order:
		r.AddOrder(v)
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}

// Добавление с сохранением
func (r *Repository) AddProduct(p model.Product) {
	r.muProducts.Lock()
	defer r.muProducts.Unlock()
	r.products = append(r.products, p)
	r.saveProducts()
}

func (r *Repository) AddCartItem(c model.CartItem) {
	r.muCart.Lock()
	defer r.muCart.Unlock()
	r.cart = append(r.cart, c)
	r.saveCart()
}

func (r *Repository) AddOrder(o model.Order) {
	r.muOrders.Lock()
	defer r.muOrders.Unlock()
	r.orders = append(r.orders, o)
	r.saveOrders()
}

// Получение копий
func (r *Repository) GetProducts() []model.Product {
	r.muProducts.Lock()
	defer r.muProducts.Unlock()
	cp := make([]model.Product, len(r.products))
	copy(cp, r.products)
	return cp
}

func (r *Repository) GetCart() []model.CartItem {
	r.muCart.Lock()
	defer r.muCart.Unlock()
	cp := make([]model.CartItem, len(r.cart))
	copy(cp, r.cart)
	return cp
}

func (r *Repository) GetOrders() []model.Order {
	r.muOrders.Lock()
	defer r.muOrders.Unlock()
	cp := make([]model.Order, len(r.orders))
	copy(cp, r.orders)
	return cp
}

// Сохранение в JSON
func (r *Repository) saveProducts() {
	file, err := os.Create(r.productsFile)
	if err != nil {
		fmt.Println("Ошибка сохранения products:", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(r.products)
}

func (r *Repository) saveCart() {
	file, err := os.Create(r.cartFile)
	if err != nil {
		fmt.Println("Ошибка сохранения cart:", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(r.cart)
}

func (r *Repository) saveOrders() {
	file, err := os.Create(r.ordersFile)
	if err != nil {
		fmt.Println("Ошибка сохранения orders:", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(r.orders)
}

// Загрузка из JSON
func (r *Repository) loadProducts() {
	file, err := os.Open(r.productsFile)
	if err != nil {
		fmt.Println("Файл products не найден, создаём новый")
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&r.products)
}

func (r *Repository) loadCart() {
	file, err := os.Open(r.cartFile)
	if err != nil {
		fmt.Println("Файл cart не найден, создаём новый")
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&r.cart)
}

func (r *Repository) loadOrders() {
	file, err := os.Open(r.ordersFile)
	if err != nil {
		fmt.Println("Файл orders не найден, создаём новый")
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&r.orders)
}
