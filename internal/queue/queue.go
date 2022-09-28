package queue

import (
	"log"

	"github.com/goocarry/wb-internship/internal/config"
	stan "github.com/nats-io/stan.go"
)

// NATS ...
type NATS struct {
	NConn  stan.Conn
	Config *config.Config
}

// New ...
func New(config *config.Config, logger *log.Logger) (*NATS, error) {
	return &NATS{
		NConn:  connect(config, logger),
		Config: config,
	}, nil
}

// Connect to NATS Streaming
func connect(config *config.Config, logger *log.Logger) stan.Conn {

	var conn stan.Conn
	var err error

	logger.Println("info-5aad4792: init NATS streaming", config.NATSClusterID, config.NATSClientID)

	conn, err = stan.Connect(config.NATSClusterID, config.NATSClientID, stan.Pings(1, 2),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			logger.Fatalf("err-feca3aa0: connection lost, reason: %v", reason)
		}))
	if err != nil {
		logger.Fatalf("err-10b56d01: can't connect: %v,\nmake sure a NATS Streaming Server is running at: %s", err, stan.DefaultNatsURL)
	}
	logger.Printf("info-45448a99: connected to %s clusterID: [%s] clientID: [%s]\n", stan.DefaultNatsURL, config.NATSClusterID, config.NATSClientID)

	return conn
}
