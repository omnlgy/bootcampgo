package seed

import (
	"log"
	"time"

	"example.com/internal/models"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {

	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderItem{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	now := time.Date(2026, 6, 23, 13, 4, 35, 583488000, time.UTC)
	now2 := time.Date(2026, 6, 23, 13, 5, 14, 778886000, time.UTC)

	// Users — unique on email
	for _, u := range []models.User{
		{Email: "alice@example.com", CreatedAt: now},
		{Email: "bob@example.com", CreatedAt: now},
		{Email: "carol@example.com", CreatedAt: now},
		{Email: "dave@example.com", CreatedAt: now2},
		{Email: "eve@example.com", CreatedAt: now2},
		{Email: "frank@example.com", CreatedAt: now2},
	} {
		db.FirstOrCreate(&u, models.User{Email: u.Email})
	}
	log.Print("seeded users")

	// Products — unique on sku
	for _, p := range []models.Product{
		{SKU: "PRO-001", Name: "Wireless Mouse", Price: 29.99, Stock: 50, CreatedAt: now},
		{SKU: "PRO-002", Name: "USB-C Hub", Price: 45.00, Stock: 30, CreatedAt: now},
		{SKU: "PRO-003", Name: "Webcam HD", Price: 89.99, Stock: 20, CreatedAt: now},
		{SKU: "PRO-004", Name: "Desk Lamp", Price: 39.99, Stock: 25, CreatedAt: now},
		{SKU: "PRO-005", Name: "Mechanical Keyboard", Price: 129.99, Stock: 15, CreatedAt: now2},
		{SKU: "PRO-006", Name: `Monitor 27"`, Price: 349.99, Stock: 10, CreatedAt: now2},
		{SKU: "PRO-007", Name: "Headset", Price: 79.99, Stock: 20, CreatedAt: now2},
		{SKU: "PRO-008", Name: "Laptop Stand", Price: 34.99, Stock: 40, CreatedAt: now2},
	} {
		db.FirstOrCreate(&p, models.Product{SKU: p.SKU})
	}
	log.Print("seeded products")

	// Orders — unique on order_number
	for _, o := range []models.Order{
		{UserID: 1, OrderNumber: "ORD-001", TotalAmount: 74.99, Status: "COMPLETED", CreatedAt: now},
		{UserID: 3, OrderNumber: "ORD-002", TotalAmount: 209.98, Status: "PENDING", CreatedAt: now2},
	} {
		db.FirstOrCreate(&o, models.Order{OrderNumber: o.OrderNumber})
	}
	log.Print("seeded orders")

	// Order items — composite PK, so FirstOrCreate uses both order_id + product_id
	for _, item := range []models.OrderItem{
		{OrderID: 1, ProductID: 1, Quantity: 1, Price: 29.99},
		{OrderID: 1, ProductID: 2, Quantity: 1, Price: 45.00},
		{OrderID: 2, ProductID: 5, Quantity: 1, Price: 129.99},
		{OrderID: 2, ProductID: 7, Quantity: 1, Price: 79.99},
	} {
		db.FirstOrCreate(&item, models.OrderItem{OrderID: item.OrderID, ProductID: item.ProductID})
	}
	log.Print("seeded order_items")

	log.Println("seed complete")
}
