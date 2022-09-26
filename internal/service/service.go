package service

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/goocarry/wb-internship/internal/config"
	"github.com/goocarry/wb-internship/internal/handlers"
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
	// err := store.Open()
	// if err != nil {
	// 	return &Service{}, err
	// }

	// defer store.Close()

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
	err := s.store.Open()
	if err != nil {
		log.Fatal(err)
	}

	queue, err := queue.New(s.cfg, s.logger)
	if err != nil {
		log.Fatal(err)
	}

	s.logger.Println("info-asd9g21d: subscribing to NATS")
	_, _ = queue.NConn.Subscribe("wbl0topic", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}, stan.DeliverAllAvailable())

	defer s.store.Close()

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
