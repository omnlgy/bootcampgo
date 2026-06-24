package service

import (
	"errors"
	"fmt"
	"time"

	"example.com/internal/domain"
	"example.com/internal/models"
	"example.com/internal/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderService struct {
	orderRepo     domain.OrderRepository
	productRepo   domain.ProductRepository
	userRepo      domain.UserRepository
	orderItemRepo domain.OrderItemRepository
	transactor    repository.TransactionRunner
}

func NewOrderService(orderRepo domain.OrderRepository, productRepo domain.ProductRepository, userRepo domain.UserRepository, orderItemRepo domain.OrderItemRepository, transactor repository.TransactionRunner) *OrderService {
	return &OrderService{
		orderRepo:     orderRepo,
		productRepo:   productRepo,
		userRepo:      userRepo,
		orderItemRepo: orderItemRepo,
		transactor:    transactor,
	}
}

func (s *OrderService) CreateOrder(order models.Order) error {
	var newOrder models.Order

	return s.transactor.WithinTransaction(func(tx *gorm.DB) error {
		orderRepo := s.orderRepo.WithTx(tx)
		productRepo := s.productRepo.WithTx(tx)
		userRepo := s.userRepo.WithTx(tx)
		// orderItemRepo := s.orderItemRepo.WithTx(tx)

		if _, err := userRepo.GetByID(order.UserID); err != nil {
			return err
		}

		var totalAmount float64
		var orderItems []models.OrderItem
		var orderNumber = generateOrderNumber()

		// Calculate totals and validate stock first
		for _, item := range order.Items {
			var product models.Product

			errLock := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, item.ProductID).Error
			if errLock != nil {
				return errLock
			}

			if product.Stock < item.Quantity {
				return fmt.Errorf("insufficient stock for product %d", item.ProductID)
			}

			newStock := product.Stock - item.Quantity
			product.Stock = newStock
			if _, err := productRepo.Update(product); err != nil {
				return err
			}

			subTotal := product.Price * float64(item.Quantity)
			totalAmount += subTotal

			orderItems = append(orderItems, models.OrderItem{
				ProductID: product.ID,
				Quantity:  item.Quantity,
				Price:     product.Price,
			})
		}

		newOrder = models.Order{
			UserID:      order.UserID,
			TotalAmount: totalAmount,
			OrderNumber: orderNumber,
			Status:      "COMPLETED",
			Items:       orderItems,
		}
		return errors.New("tes error")
		// Create order first — GORM backfills the ID and inserts items via the association
		if err := orderRepo.Create(newOrder); err != nil {
			return err
		}

		// for _, item := range orderItems {
		// 	if err := orderItemRepo.Create(item); err != nil {
		// 		return err
		// 	}
		// }

		return nil
	})
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.orderRepo.GetAll()
}

func (s *OrderService) GetOrderByID(id uint) (models.Order, error) {
	return s.orderRepo.GetByID(id)
}

func (s *OrderService) UpdateOrder(order models.Order) (models.Order, error) {
	return s.orderRepo.Update(order)
}

func (s *OrderService) DeleteOrder(id uint) error {
	return s.orderRepo.Delete(id)
}

func generateOrderNumber() string {
	return fmt.Sprintf("INV/%s/%d", time.Now().Format("20060102"), time.Now().UnixNano()%10000)
}
