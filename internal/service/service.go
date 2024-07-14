package service

import (
	"rest-api-example/internal/models"
	"rest-api-example/internal/repository"
)

type UserServiceList interface {
	Create(user models.User) (int, error)
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	Update(id int, input models.UpdateUser) error
	Delete(id int) error
}

type ProductServiceList interface {
	GetAllProducts() ([]models.Product, error)
	GetProductsById(id int) (models.Product, error)
}

type OrderServiceList interface {
	CreateOrder(order models.Order) (int, error)
	GetAllOrders(userId int) ([]models.Order, error)
}

type Service struct {
	UserServiceList
	ProductServiceList
	OrderServiceList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserServiceList:    NewUserService(repos),
		ProductServiceList: NewProductService(repos),
		OrderServiceList:   NewOrderService(repos),
	}
}
