package dto

type BodyProduct struct {
	SKU       string  `json:"sku" binding:"required"`
	Name      string  `json:"name" binding:"required"`
	Price     float64 `json:"price" binding:"required,gt=0"`
	Stock     int     `json:"stock" binding:"gte=0"`
	ExpiredAt string  `json:"expired_at" binding:"required,expire_range"`
}
