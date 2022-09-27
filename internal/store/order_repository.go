package store

import (
	"encoding/json"

	"github.com/goocarry/wb-internship/internal/model"
)

// OrderRepository ...
type OrderRepository struct {
	store *Store
}

// Create ...
func (r *OrderRepository) Create(o *model.Order) (*model.Order, error) {
	ordersrt, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	err = r.store.db.QueryRow(`INSERT INTO public."order_cache" (order_uid, data) VALUES ($1, $2) RETURNING order_uid`,
		o.OrderUID,
		ordersrt,
	).Scan(&o.OrderUID)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// GetAll ...
func (r *OrderRepository) GetAll() (map[string]model.Order, error) {
	rows, err := r.store.db.Query(`SELECT order_uid, data FROM public."order_cache"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make(map[string]model.Order)

	for rows.Next() {
		ordercache := model.OrderCache{}
		order := model.Order{}

		if err := rows.Scan(&ordercache.OrderUID, &ordercache.Order); err != nil {
			return orders, err
		}
		json.Unmarshal([]byte(ordercache.Order), &order)
		orders[order.OrderUID] = order
	}

	return orders, nil
}
