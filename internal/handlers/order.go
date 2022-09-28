package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/goocarry/wb-internship/internal/cache"
	"github.com/goocarry/wb-internship/internal/model"
	"github.com/goocarry/wb-internship/internal/store"
	"github.com/gorilla/mux"
)

// Order ...
type Order struct {
	logger *log.Logger
	store  *store.Store
	cache  *cache.Cache
}

// NewOrder ...
func NewOrder(l *log.Logger, s *store.Store, c *cache.Cache) *Order {
	return &Order{
		logger: l,
		store:  s,
		cache:  c,
	}
}

// CreateOrder ...
func (o *Order) CreateOrder(rw http.ResponseWriter, r *http.Request) {
	o.logger.Println("info-b23e7145: creating new order")

	order := &model.Order{}
	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		o.logger.Printf("err-44344769: can't parse data: %s", err)
		http.Error(rw, "can't parse data", http.StatusBadRequest)
		return
	}

	_, err := o.store.Order().Create(order)
	if err != nil {
		o.logger.Printf("err-b9e75800: can't create order: %s", err)
		http.Error(rw, "can't create order", http.StatusBadRequest)
		return
	}

	o.logger.Printf("info-f8a835a5: new order created: %v", order)
	fmt.Fprint(rw, "new order created", http.StatusOK)
}

// GetOrder ...
func (o *Order) GetOrder(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	o.logger.Printf("info-b23e7145: get order: %s", params["id"])
	order, exists := o.cache.Get(params["id"])
	if exists {
		fmt.Fprintf(rw, "order: %v", order)
	} else {
		fmt.Fprint(rw, "order not found", http.StatusOK)
	}
}
