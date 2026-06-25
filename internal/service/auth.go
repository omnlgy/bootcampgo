package service

import (
	"time"

	"example.com/internal/domain"
	"example.com/internal/models"
	"example.com/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", ValidationError{
			Field:   "username,password",
			Rule:    "invalid",
			Message: "invalid credentials",
		}
	}

	if !utils.CheckPasswordHash(user.Password, password) {
		return "", ValidationError{
			Field:   "username,password",
			Rule:    "invalid",
			Message: "invalid credentials",
		}
	}

	token, err := GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Register(email, password string) (string, error) {
	if _, err := s.userRepo.GetByEmail(email); err == nil {
		return "", &ConflictError{
			Resource: "user",
			Field:    "email",
			value:    email,
		}
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	user := models.User{
		Email:    email,
		Password: hashed,
	}

	if err := s.userRepo.Create(user); err != nil {
		return "", err
	}

	token, err := GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

const secretKey = "supersecretprivatekey"

func GenerateJWT(userId uint, email string) (string, error) {
	claims := JWTClaims{
		UserID: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, ValidationError{
			Field:   "token",
			Rule:    "invalid",
			Message: "invalid token",
		}
	}

	return claims, nil
}
