package service

import (
	"go-fiber-clean-api/internals/user/entity"
	"go-fiber-clean-api/internals/user/repository"
)

type UserService interface {
	GetUsers() ([]entity.User, error)
	GetUser(id uint) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUsers() ([]entity.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUser(id uint) (entity.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) CreateUser(user entity.User) (entity.User, error) {
	return s.repo.Create(user)
}

func (s *userService) UpdateUser(user entity.User) (entity.User, error) {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
