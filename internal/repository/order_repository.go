package repository

import (
	"gorm.io/gorm"
	"log"
	"wb_tech_level_0/pkg/model"
)

// Репозиторий для работы с заказами

type OrderRepository struct {
	db *gorm.DB
}

// Конструктор

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Сохраняет заказ в базе данных

func (r *OrderRepository) Save(order *model.Order) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Сначала сохраняем Delivery
		if err := tx.Create(&order.Delivery).Error; err != nil {
			log.Printf("failed to save delivery: %v", err)
			return err
		}

		// Далее сохраняем Payment
		if err := tx.Create(&order.Payment).Error; err != nil {
			log.Printf("failed to save payment: %v", err)
			return err
		}

		// Присваиваем сохраненные DeliveryID и PaymentID в Order
		order.DeliveryID = order.Delivery.ID
		order.PaymentID = order.Payment.ID

		// Сохраняем Order
		if err := tx.Create(&order).Error; err != nil {
			log.Printf("failed to save order: %v", err)
			return err
		}

		// Теперь сохраняем каждый Item и связываем с Order
		for i := range order.Items {
			order.Items[i].OrderUID = order.OrderUID

			// Проверяем, существует ли уже запись с таким же ID
			var existingItem model.Item
			if err := tx.First(&existingItem, "order_uid = ? AND chrt_id = ?", order.OrderUID, order.Items[i].ChrtID).Error; err == nil {
				// Если запись существует, обновляем её
				if err := tx.Model(&existingItem).Updates(order.Items[i]).Error; err != nil {
					log.Printf("failed to update order item: %v", err)
					return err
				}
			} else if err := tx.Create(&order.Items[i]).Error; err != nil {
				log.Printf("failed to save order item: %v", err)
				return err
			}
		}

		return nil
	})
}

// Извлекает все заказы

func (r *OrderRepository) GetAll() ([]model.Order, error) {
	var orders []model.Order
	if err := r.db.Preload("Delivery").Preload("Payment").Preload("Items").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
