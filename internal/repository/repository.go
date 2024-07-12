package repository

import (
	"database/sql"
	"rest-api-example/internal/models"
)

type UserRepository interface {
	Create(user models.User) (int, error)
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	Update(id int, input models.UpdateUser) error
	Delete(id int) error
}

type Repository struct {
	UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: NewUserPostgres(db),
	}
}
