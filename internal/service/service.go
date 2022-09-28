package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"net/http"

	"github.com/goocarry/wb-internship/internal/cache"
	"github.com/goocarry/wb-internship/internal/config"
	"github.com/goocarry/wb-internship/internal/handlers"
	"github.com/goocarry/wb-internship/internal/model"
	"github.com/goocarry/wb-internship/internal/queue"
	"github.com/goocarry/wb-internship/internal/store"
	"github.com/gorilla/mux"
	stan "github.com/nats-io/stan.go"
)

// Service ...
type Service struct {
	cfg    *config.Config
	logger *log.Logger
	store  *store.Store
	cache  *cache.Cache
	router *mux.Router
}

// NewService ...
func NewService(config *config.Config, logger *log.Logger) (*Service, error) {

	logger.Println("info-8e0b1b3b: creating new service")
	// create the store
	store := store.New(config)

	// create the cache
	cache := cache.New()

	// create the router
	router := mux.NewRouter()

	service := &Service{
		cfg:    config,
		logger: logger,
		store:  store,
		cache:  cache,
		router: router,
	}

	service.configureRoutes(logger, store, cache)

	return service, nil
}

// Run ...
func (s *Service) Run() error {
	// open db connection
	err := s.store.Open()
	if err != nil {
		log.Fatal(err)
	}

	// dborders, err := s.store.Order().GetAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// s.cache.SetAll(dborders)

	// init nats connection
	queue, err := queue.New(s.cfg, s.logger)
	if err != nil {
		log.Fatal(err)
	}

	s.logger.Println("info-asd9g21d: subscribing to NATS")
	_, _ = queue.NConn.Subscribe("wbl0topic", s.onMessage, stan.StartWithLastReceived())

	defer s.store.Close()
	defer queue.NConn.Close()

	return s.startHTTP()
}

// configureRoutes ...
func (s *Service) configureRoutes(logger *log.Logger, store *store.Store, cache *cache.Cache) {
	// create the handlers
	orderHandler := handlers.NewOrder(logger, store, cache)

	s.router.HandleFunc("/order", orderHandler.CreateOrder).Methods("POST")
	s.router.HandleFunc("/order/{id}", orderHandler.GetOrder)
}

// startHTTP ...
func (s *Service) startHTTP() error {
	s.logger.Println("info-b0f45365: start HTTP")

	var listener net.Listener

	s.logger.Printf("info-af3aaf35: bind application to host: %s and port: %s", s.cfg.Listen.BindIP, s.cfg.Listen.Port)
	var err error
	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", s.cfg.Listen.BindIP, s.cfg.Listen.Port))
	if err != nil {
		s.logger.Fatal(err)
	}

	s.logger.Println("info-9faad0c7: application completely initialized")

	return http.Serve(listener, s.router)
}

func (s *Service) onMessage(m *stan.Msg) {
	newOrder := &model.Order{}
	s.logger.Printf("info-df0d62b3: received a message: %s\n", string(m.Data))
	err := json.Unmarshal([]byte(m.Data), newOrder)
	if err != nil || newOrder.OrderUID == "" {
		s.logger.Printf("err-e4b02fb1: order data is not valid")
		return
	}

	_, exist := s.cache.Get(newOrder.OrderUID)
	if exist != false {
		s.logger.Printf("info-f9358936: order with this order_uid already exists: %s\n", newOrder.OrderUID)
		return
	}

	// add neworder to cache
	s.cache.Set(newOrder.OrderUID, newOrder)

	// save neworder to pg
	_, err = s.store.Order().Create(newOrder)
	if err != nil {
		s.logger.Fatal(err)
	}
}
