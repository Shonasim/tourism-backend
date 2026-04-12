package service

import (
	"errors"
	"tourism-backend/internal/domain"
	"tourism-backend/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(name, email, password string) (*domain.User, error) {
	exist, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return nil, errors.New("user with this email already exists")
	}
	user := &domain.User{
		Name:         name,
		Email:        email,
		PasswordHash: password,
		Role:         domain.RoleClient,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetByID(id int) (*domain.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) GetAll() ([]*domain.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
