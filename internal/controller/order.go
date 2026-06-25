package controller

import (
	"example.com/internal/dto"
	"example.com/internal/models"
	"example.com/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var body dto.BodyCreateOrder
	if err := ctx.ShouldBindJSON(&body); err != nil {
		mapError(ctx, err)
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
		mapError(ctx, err)
		return
	}

	ctx.JSON(201, gin.H{"message": "order created"})
}

func (oc *OrderController) GetOrders(ctx *gin.Context) {
	orders, err := oc.orderService.GetAllOrders()
	if err != nil {
		mapError(ctx, err)
		return
	}
	ctx.JSON(200, orders)
}
