package service

import (
	"rest-api-example/internal/models"
	"rest-api-example/internal/repository"
)

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order models.Order) (int, error) {
	return s.repo.CreateOrder(order)
}

func (s *OrderService) GetAllOrders(userId int) ([]models.Order, error) {
	return s.repo.GetAllOrders(userId)
}
