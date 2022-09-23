package store

import (
	"database/sql"

	"github.com/goocarry/wb-internship/internal/config"
	//
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config          *config.Config
	db              *sql.DB
	orderRepository *OrderRepository
}

// New ...
func New(config *config.Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	dsn := s.config.PostgresURL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() error {
	return s.db.Close()
}

// Order ...
func (s *Store) Order() *OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}

	s.orderRepository = &OrderRepository{
		store: s,
	}

	return s.orderRepository
}
