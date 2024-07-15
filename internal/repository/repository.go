package repository

import (
	"database/sql"
	"rest-api-example/internal/models"
)

type UserRepository interface {
	Create(user models.User) (int, error)
	GetAll() ([]models.UserInfo, error)
	GetById(id int) (models.UserInfo, error)
	Update(id int, input models.UpdateUser) error
	Delete(id int) error
}

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductsById(id int) (models.Product, error)
}

type OrderRepository interface {
	CreateOrder(order models.Order) (int, error)
	GetAllOrders(userId int) ([]models.Order, error)
}

type Repository struct {
	UserRepository
	ProductRepository
	OrderRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository:    NewUserPostgres(db),
		ProductRepository: NewProductPostgres(db),
		OrderRepository:   NewOrderPostgres(db),
	}
}
