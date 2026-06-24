package models

import "time"

type Order struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	UserID      uint        `json:"user_id"`
	User        User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	OrderNumber string      `gorm:"type:varchar(50);uniqueIndex;not null" json:"order_number"`
	TotalAmount float64     `gorm:"type:numeric(12,2)" json:"total_amount"`
	Status      string      `gorm:"type:varchar(20);default:'COMPLETED'" json:"status"`
	Items       []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt   time.Time   `json:"created_at"`
}

func (Order) TableName() string {
	return "orders"
}
