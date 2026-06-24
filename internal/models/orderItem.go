package models

type OrderItem struct {
	OrderID   uint    `gorm:"primaryKey" json:"order_id"`
	ProductID uint    `gorm:"primaryKey" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int     `gorm:"type:int;not null" json:"quantity"`
	Price     float64 `gorm:"type:numeric(12,2);not null" json:"price"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
