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

type Service struct {
	UserServiceList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserServiceList: NewUserService(repos),
	}
}
