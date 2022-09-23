package store

import "github.com/goocarry/wb-internship/internal/model"

// OrderRepository ...
type OrderRepository struct {
	store *Store
}

// Create ...
func (r *OrderRepository) Create(o *model.Order) (*model.Order, error) {
	err := r.store.db.QueryRow(`INSERT INTO public."order" (order_uid, track_number, entry) VALUES ($1, $2, $3) RETURNING order_uid`,
		o.OrderUID,
		o.TrackNumber,
		o.Entry,
	).Scan(&o.OrderUID)
	if err != nil {
		return nil, err
	}
	return o, nil
}
