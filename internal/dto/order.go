package dto

type OrderItemInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
}

type BodyCreateOrder struct {
	UserID uint             `json:"user_id" binding:"required"`
	Items  []OrderItemInput `json:"items" binding:"required,dive"`
}
