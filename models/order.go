package models

import (
	"time"

	// "gorm.io/gorm"
)

type Order struct {
	ID	   uint      `json:"id" gorm:"primaryKey"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	OrderDate time.Time `json:"order_date"`
}