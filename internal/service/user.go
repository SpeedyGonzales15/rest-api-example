package service

import (
	"rest-api-example/internal/models"
	"rest-api-example/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user models.User) (int, error) {
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(id int) (models.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Update(id int, input models.UpdateUser) error {
	return s.repo.Update(id, input)
}

func (s *UserService) Delete(id int) error {
	return s.repo.Delete(id)
}
