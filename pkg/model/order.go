package model

import (
	"gorm.io/gorm"
	"time"
)

// Order модель для заказа
type Order struct {
	ID                uint   `gorm:"primaryKey;autoIncrement"`
	OrderUID          string `gorm:"unique"`
	TrackNumber       string
	Entry             string
	DeliveryID        uint
	Delivery          Delivery
	PaymentID         uint
	Payment           Payment
	Locale            string
	InternalSignature string
	CustomerID        string
	DeliveryService   string
	Shardkey          string
	SmID              int
	DateCreated       time.Time
	OofShard          string
	Items             []Item `gorm:"foreignKey:OrderUID;references:OrderUID"`
}

// Delivery модель для доставки
type Delivery struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Phone     string
	Zip       string
	City      string
	Address   string
	Region    string
	Email     string
}

// Payment модель для оплаты
type Payment struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Transaction  string
	RequestID    string
	Currency     string
	Provider     string
	Amount       int
	PaymentDT    int64
	Bank         string
	DeliveryCost int
	GoodsTotal   int
	CustomFee    int
}

// Item модель для элемента заказа
type Item struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	ChrtID      int
	TrackNumber string
	Price       int
	Rid         string
	Name        string
	Sale        int
	Size        string
	TotalPrice  int
	NmID        int
	Brand       string
	Status      int
	OrderUID    string `gorm:"index"`
}
