package controller

import (
	"example.com/internal/models"
	"example.com/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *service.OrderService
}

type OrderItemInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
}

type CheckoutInput struct {
	UserID uint             `json:"user_id" binding:"required"`
	Items  []OrderItemInput `json:"items" binding:"required,dive"`
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var body CheckoutInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var orderItems []models.OrderItem
	for _, item := range body.Items {
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}
	if err := oc.orderService.CreateOrder(models.Order{
		UserID: body.UserID,
		Items:  orderItems,
	}); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "order created"})
}
