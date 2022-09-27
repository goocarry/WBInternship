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
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	stan "github.com/nats-io/stan.go"
)

// Service ...
type Service struct {
	cfg    *config.Config
	logger *log.Logger
	store  *store.Store
	router *mux.Router
}

// NewService ...
func NewService(config *config.Config, logger *log.Logger) (*Service, error) {

	// create the store
	store := store.New(config)

	// create the router
	router := mux.NewRouter()

	service := &Service{
		cfg:    config,
		logger: logger,
		store:  store,
		router: router,
	}

	service.configureRoutes(logger, store)

	return service, nil
}

// Run ...
func (s *Service) Run() error {
	// open db connection
	err := s.store.Open()
	if err != nil {
		log.Fatal(err)
	}

	// new cache instance
	cache := cache.New()
	dborders, err := s.store.Order().GetAll()
	if err != nil {
		log.Fatal(err)
	}
	cache.SetAll(dborders)

	// init nats connection
	queue, err := queue.New(s.cfg, s.logger)
	if err != nil {
		log.Fatal(err)
	}

	s.logger.Println("info-asd9g21d: subscribing to NATS")
	_, _ = queue.NConn.Subscribe("wbl0topic", func(m *stan.Msg) {
		newOrder := &model.Order{}
		fmt.Printf("Received a message: %s\n", string(m.Data))
		json.Unmarshal([]byte(m.Data), newOrder)

		_, exist := cache.Get(newOrder.OrderUID)
		if exist != false {
			log.Printf("order with this order_uid already exists: %s\n", newOrder.OrderUID)
			return
		}

		// add neworder to cache
		cache.Set(newOrder.OrderUID, newOrder)

		// save neworder to pg
		_, err := s.store.Order().Create(newOrder)
		if err != nil {
			log.Fatal(err)
		}

	}, stan.StartWithLastReceived())

	defer s.store.Close()
	defer queue.NConn.Close()

	return s.startHTTP()
}

// configureRoutes ...
func (s *Service) configureRoutes(logger *log.Logger, store *store.Store) {
	// create the handlers
	helloHandler := handlers.NewHello(logger)
	orderHandler := handlers.NewOrder(logger, store)

	s.router.Use(gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/", helloHandler.Hello)
	s.router.HandleFunc("/order", orderHandler.CreateOrder).Methods("POST")
}

// startHTTP ...
func (s *Service) startHTTP() error {
	s.logger.Println("Start HTTP")

	var listener net.Listener

	s.logger.Printf("Bind application to host: %s and port: %s", s.cfg.Listen.BindIP, s.cfg.Listen.Port)
	var err error
	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", s.cfg.Listen.BindIP, s.cfg.Listen.Port))
	if err != nil {
		s.logger.Fatal(err)
	}

	s.logger.Println("Application completely initialized and started")

	return http.Serve(listener, s.router)
}
