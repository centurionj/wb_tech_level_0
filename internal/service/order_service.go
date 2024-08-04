package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"wb_tech_level_0/internal/repository"
	"wb_tech_level_0/pkg/model"
)

type OrderService struct {
	orderRepo   *repository.OrderRepository
	redisClient *redis.Client
}

func NewOrderService(orderRepo *repository.OrderRepository, redisClient *redis.Client) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		redisClient: redisClient,
	}
}

// SaveOrder сохраняет заказ в Redis и в базе данных
func (s *OrderService) SaveOrder(order *model.Order) error {
	ctx := context.Background()

	// Сохраняем заказ в базе данных
	if err := s.orderRepo.Save(order); err != nil {
		return err
	}

	// Сохраняем заказ в Redis
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	if err := s.redisClient.Set(ctx, string(order.ID), orderJSON, 0).Err(); err != nil {
		return err
	}

	return nil
}

// GetOrder извлекает заказ из Redis
func (s *OrderService) GetOrder(orderID uint) (*model.Order, error) {
	ctx := context.Background()

	// Пытаемся получить заказ из Redis
	orderJSONStr, err := s.redisClient.Get(ctx, string(orderID)).Result()
	if err == redis.Nil {
		// Данные отсутствуют в Redis
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var order model.Order
	if err := json.Unmarshal([]byte(orderJSONStr), &order); err != nil {
		return nil, err
	}

	return &order, nil
}

// Pаполняет Redis данными из базы данных
func (s *OrderService) FillCache() error {
	ctx := context.Background()

	// Извлекаем все заказы из базы данных
	orders, err := s.orderRepo.GetAll()
	if err != nil {
		return err
	}

	// Сохраняем каждый заказ в Redis
	for _, order := range orders {
		orderJSON, err := json.Marshal(order)
		if err != nil {
			return err
		}
		if err := s.redisClient.Set(ctx, string(order.ID), orderJSON, 0).Err(); err != nil {
			return err
		}
	}

	return nil
}
