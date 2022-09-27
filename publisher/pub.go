package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/goocarry/wb-internship/internal/model"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

func main() {

	file, err := ioutil.ReadFile("model.json")
	if err != nil {
		log.Fatal(err)
	}
	order := model.Order{}

	json.Unmarshal([]byte(file), &order)
	order.OrderUID = uuid.New().String()
	updatedOrder, _ := json.Marshal(order)

	fmt.Println("Order Id: ", order.OrderUID)

	sc, err := stan.Connect("test-cluster", "wbl0natspublisher")
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan.DefaultNatsURL)
	}
	defer sc.Close()

	err = sc.Publish("wbl0topic", updatedOrder)
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
	log.Printf("Published [%s] : '%s'\n", "wbl0topic", order.OrderUID)

}
