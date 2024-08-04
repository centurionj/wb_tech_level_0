package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"wb_tech_level_0/pkg/model"
)

func main() {
	sc, err := stan.Connect("test-cluster", "publisher-client", stan.NatsURL(
		fmt.Sprintf("nats://%s:%s", os.Getenv("NATS_HOST"), os.Getenv("NATS_PORT"))))
	if err != nil {
		log.Fatalf("failed to connect to nats-streaming: %v", err)
	}
	defer sc.Close()

	order := model.Order{
		OrderUID:    "b563feb7b2b84b6test_1",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: model.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: model.Payment{
			Transaction:  "b563feb7b2b84b6test",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDT:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
		},
		Items: []model.Item{
			{
				ChrtID:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				Rid:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NmID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
	}

	// Сериализация заказа в JSON
	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("failed to marshal order: %v", err)
	}

	// Публикация сообщения в NATS Streaming
	if err := sc.Publish("orders", orderJSON); err != nil {
		log.Fatalf("failed to publish order: %v", err)
	}

	log.Println("Order published successfully")
}
