package main

import (
	"log"
	"wb_tech_level_0/config"
	"wb_tech_level_0/internal/cache"
	"wb_tech_level_0/internal/database"
	"wb_tech_level_0/internal/nats"
	"wb_tech_level_0/internal/repository"
	"wb_tech_level_0/internal/server"
	"wb_tech_level_0/internal/service"
)

func main() {
	// Загрузка конфига
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Подключение к дб
	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	// Подключение к Redis
	redisClient, err := cache.NewRedisClient(cfg)
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, redisClient)

	// Заполнение кэша данными из базы данных
	if err := orderService.FillCache(); err != nil {
		log.Fatalf("failed to fill cache: %v", err)
	}

	// Запуск NATS подписки в горутине
	go nats.Subscribe(cfg, orderService)

	// Запуск HTTP сервера
	httpServer := server.NewHTTPServer(cfg, orderService)
	if err := httpServer.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
