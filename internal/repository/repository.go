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

	LoadedProducts int
	LoadedCart     int
	LoadedOrders   int
}

// NewRepository создаёт репозиторий и загружает данные из файлов.
// Передавайте пути к JSON-файлам для каждого слайса.
func NewRepository(productsFile, cartFile, ordersFile string) *Repository {
	r := &Repository{
		productsFile: productsFile,
		cartFile:     cartFile,
		ordersFile:   ordersFile,
	}

	// загружаем данные
	if err := r.loadProducts(); err != nil {
		fmt.Printf("loadProducts error: %v\n", err)
	}
	if err := r.loadCart(); err != nil {
		fmt.Printf("loadCart error: %v\n", err)
	}
	if err := r.loadOrders(); err != nil {
		fmt.Printf("loadOrders error: %v\n", err)
	}

	// фиксируем, какие объекты были восстановлены из файла при старте
	r.LoadedProducts = len(r.products)
	r.LoadedCart = len(r.cart)
	r.LoadedOrders = len(r.orders)

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
	if err := r.saveProducts(); err != nil {
		fmt.Printf("saveProducts error: %v\n", err)
	}
}

func (r *Repository) AddCartItem(c model.CartItem) {
	r.muCart.Lock()
	defer r.muCart.Unlock()
	r.cart = append(r.cart, c)
	if err := r.saveCart(); err != nil {
		fmt.Printf("saveCart error: %v\n", err)
	}
}

func (r *Repository) AddOrder(o model.Order) {
	r.muOrders.Lock()
	defer r.muOrders.Unlock()
	r.orders = append(r.orders, o)
	if err := r.saveOrders(); err != nil {
		fmt.Printf("saveOrders error: %v\n", err)
	}
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
func (r *Repository) saveProducts() error {
	tmp := r.productsFile + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(r.products); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return os.Rename(tmp, r.productsFile)
}

func (r *Repository) saveCart() error {
	tmp := r.cartFile + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(r.cart); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return os.Rename(tmp, r.cartFile)
}

func (r *Repository) saveOrders() error {
	tmp := r.ordersFile + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(r.orders); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return os.Rename(tmp, r.ordersFile)
}

// Загрузка из JSON
func (r *Repository) loadProducts() error {
	file, err := os.Open(r.productsFile)
	if err != nil {
		if os.IsNotExist(err) {
			// файл не существует — это нормально
			return nil
		}
		return err
	}
	defer file.Close()

	var data []model.Product
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		// если файл пустой или повреждён — возвращаем ошибку,
		// но не падаем жестко (чтобы можно было отладить)
		return fmt.Errorf("decode products: %w", err)
	}

	r.muProducts.Lock()
	r.products = data
	r.muProducts.Unlock()
	return nil
}

func (r *Repository) loadCart() error {
	file, err := os.Open(r.cartFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	var data []model.CartItem
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("decode cart: %w", err)
	}

	r.muCart.Lock()
	r.cart = data
	r.muCart.Unlock()
	return nil
}

func (r *Repository) loadOrders() error {
	file, err := os.Open(r.ordersFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	var data []model.Order
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("decode orders: %w", err)
	}

	r.muOrders.Lock()
	r.orders = data
	r.muOrders.Unlock()
	return nil
}
