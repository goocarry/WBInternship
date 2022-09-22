package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/goocarry/wb-internship/internal/config"
	"github.com/goocarry/wb-internship/internal/handlers"
	"github.com/goocarry/wb-internship/pkg/client/postgresql"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Service ...
type Service struct {
	cfg        *config.Config
	logger     *log.Logger
	httpServer *http.Server
	pgClient   *pgxpool.Pool
}

// NewService ...
func NewService(config *config.Config, logger *log.Logger) (Service, error) {

	logger.Println("creating pg conifg")
	pgConfig := postgresql.NewPgConfig(config.PostgreSQL.Username, config.PostgreSQL.Password, config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database)
	pgClient, err := postgresql.NewClient(context.Background(), 5, 5*time.Second, pgConfig)
	if err != nil {
		log.Fatalf("error during config initialization")
	}

	// create the handlers
	helloHandler := handlers.NewHello(logger)

	// create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)

	// create a new server
	s := http.Server{
		Addr:         ":9990",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	return Service{
		cfg:        config,
		logger:     logger,
		pgClient:   pgClient,
		httpServer: &s,
	}, nil
}

// Run ...
func (s *Service) Run() {
	s.startHTTP()
}

// startHTTP ...
func (s *Service) startHTTP() {
	s.logger.Println("Start HTTP")

	var listener net.Listener

	s.logger.Printf("bind application to host: %s and port: %s", s.cfg.Listen.BindIP, s.cfg.Listen.Port)
	var err error
	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", s.cfg.Listen.BindIP, s.cfg.Listen.Port))
	if err != nil {
		s.logger.Fatal(err)
	}

	s.logger.Println("application completely initialized and started")

	if err := s.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			s.logger.Println("server shutdown")
		default:
			s.logger.Fatal(err)
		}
	}
	err = s.httpServer.Shutdown(context.Background())
	if err != nil {
		s.logger.Fatal(err)
	}
}
