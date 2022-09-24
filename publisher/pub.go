package main

import (
	"log"

	"github.com/nats-io/stan.go"
)

func main() {

	sc, err := stan.Connect("test-cluster", "wbl0natspublisher")
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan.DefaultNatsURL)
	}
	defer sc.Close()

	err = sc.Publish("foo", []byte("PUBLISHER"))
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
	log.Printf("Published [%s] : '%s'\n", "foo", []byte("PUBLISHER"))

}
