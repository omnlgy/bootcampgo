package main

import (
	"log"

	"example.com/cmd/seed"
	"example.com/internal/controller"
	"example.com/internal/repository"
	"example.com/internal/router"
	"example.com/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=admin password=admin123 dbname=day6 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	seed.Init(db)
	transactor := repository.NewTransactor(db)

	orderRepo := repository.NewOrderRepository(db)
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)
	orderItemRepo := repository.NewOrderItemRepository(db)

	orderSvc := service.NewOrderService(orderRepo, productRepo, userRepo, orderItemRepo, transactor)
	orderCtrl := controller.NewOrderController(orderSvc)

	server := gin.Default()

	router.OrderRoutes(server, orderCtrl)
	server.Run(":8080")
}
