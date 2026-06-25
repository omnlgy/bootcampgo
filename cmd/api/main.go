package main

import (
	"log"

	"example.com/cmd/seed"
	"example.com/internal/controller"
	"example.com/internal/repository"
	"example.com/internal/router"
	"example.com/internal/service"
	"example.com/internal/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	goplayvalidator "github.com/go-playground/validator/v10"
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

	// Register custom validators
	if v, ok := binding.Validator.Engine().(*goplayvalidator.Validate); ok {
		if err := validator.RegisterCustomValidators(v); err != nil {
			log.Fatalf("failed to register custom validators: %v", err)
		}
	}

	// Repositories
	orderRepo := repository.NewOrderRepository(db)
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)
	orderItemRepo := repository.NewOrderItemRepository(db)

	// Services
	orderSvc := service.NewOrderService(orderRepo, productRepo, userRepo, orderItemRepo, transactor)
	productSvc := service.NewProductService(productRepo)
	authSvc := service.NewAuthService(userRepo)

	// Controllers
	orderCtrl := controller.NewOrderController(orderSvc)
	productCtrl := controller.NewProductController(productSvc)
	authCtrl := controller.NewAuthController(authSvc)

	server := gin.Default()

	router.RegisterRoutes(server, orderCtrl, productCtrl)
	router.RegisterAuthRoutes(server, authCtrl)
	server.Run(":8080")
}
