package controller

import (
	"time"

	"example.com/internal/dto"
	"example.com/internal/models"
	"example.com/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var body dto.BodyProduct
	if err := ctx.ShouldBindJSON(&body); err != nil {
		mapError(ctx, err)
		return
	}

	expiredAt, err := time.Parse("2006-01-02", body.ExpiredAt)
	if err != nil {
		mapError(ctx, err)
		return
	}

	product := models.Product{
		SKU:       body.SKU,
		Name:      body.Name,
		Price:     body.Price,
		Stock:     body.Stock,
		ExpiredAt: &expiredAt,
	}

	if err := pc.productService.CreateProduct(product); err != nil {
		mapError(ctx, err)
		return
	}

	ctx.JSON(201, gin.H{"message": "product created"})
}
