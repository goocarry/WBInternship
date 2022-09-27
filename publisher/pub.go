package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/goocarry/wb-internship/internal/model"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

func main() {

	pwd, _ := os.Getwd()
	file, err := ioutil.ReadFile(pwd + "/publisher/model.json")
	if err != nil {
		log.Fatal(err)
	}
	order := model.Order{}

	json.Unmarshal([]byte(file), &order)
	order.OrderUID = uuid.New().String()
	updatedOrder, _ := json.Marshal(order)

	sc, err := stan.Connect("test-cluster", uuid.New().String())
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan.DefaultNatsURL)
	}
	defer sc.Close()

	err = sc.Publish("wbl0topic", updatedOrder)
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
	// log.Printf("Published [%s] : '%s'\n", "wbl0topic", order.OrderUID)

}
