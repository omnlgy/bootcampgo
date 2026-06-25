package controller

import (
	"net/http"

	"example.com/internal/service"
	"github.com/gin-gonic/gin"
)

func mapError(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case *service.NotFoundError:
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": e.Error(),
		})
	case *service.ConflictError:
		ctx.JSON(http.StatusConflict, gin.H{
			"code":    http.StatusConflict,
			"message": e.Error(),
		})
	case *service.ValidationError:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "reques validation failed",
			"details": e,
		})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "an unexpected error occurred",
		})
	}
}
