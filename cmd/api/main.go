package main

import (
	"log"

	"example.com/internal/router"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=admin password=admin123 dbname=day6 port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// Auto-migrate models
	// err = db.AutoMigrate(
	// 	&models.User{},
	// 	&models.Concert{},
	// 	&models.Booking{},
	// 	&models.TicketCategory{},
	// 	&models.TicketDetail{},
	// )
	// if err != nil {
	// 	log.Fatalf("failed to migrate: %v", err)
	// }

	server := gin.Default()

	router.RegisteredRoute(server)
	server.Run(":8080")
}
