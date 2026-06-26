package controller

import (
	"strings"

	"example.com/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) LoginHandler(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		mapError(ctx, err)
		return
	}

	token, err := ac.authService.Login(req.Email, req.Password)
	if err != nil {
		mapError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (ac *AuthController) RegisterHandler(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		mapError(ctx, err)
		return
	}

	token, err := ac.authService.Register(req.Email, req.Password)
	if err != nil {
		mapError(ctx, err)
		return
	}

	ctx.JSON(201, gin.H{
		"token": token,
	})
}

func (ac *AuthController) LogoutHandler(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(401, gin.H{
			"code":    401,
			"message": "missing authorization header",
		})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		ctx.JSON(401, gin.H{
			"code":    401,
			"message": "invalid authorization format",
		})
		return
	}

	if err := ac.authService.Logout(tokenString); err != nil {
		mapError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "logged out successfully",
	})
}
