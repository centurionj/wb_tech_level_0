package nats

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"wb_tech_level_0/config"
	"wb_tech_level_0/internal/service"
	"wb_tech_level_0/pkg/model"
)

// Subscribe подключается к NATS-Streaming и обрабатывает сообщения
func Subscribe(cfg *config.Config, orderService *service.OrderService) {
	sc, err := stan.Connect("test-cluster", "subscriber-client", stan.NatsURL(cfg.NatsURL))
	if err != nil {
		log.Fatalf("failed to connect to nats-streaming: %v", err)
	}
	defer sc.Close()

	_, err = sc.Subscribe("orders", func(m *stan.Msg) {
		var order model.Order
		if err := json.Unmarshal(m.Data, &order); err != nil {
			log.Printf("invalid message format: %v", err)
			return
		}

		if err := orderService.SaveOrder(&order); err != nil {
			log.Printf("failed to save order: %v", err)
		}
	}, stan.DeliverAllAvailable())

	if err != nil {
		log.Fatalf("failed to subscribe to orders: %v", err)
	}

	select {}
}
