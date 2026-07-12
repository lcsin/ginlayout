package service

import (
	"context"
	"errors"
	"ginlayout/internal/model"
	"ginlayout/internal/repository"
)

type UserService interface {
	Register(ctx context.Context, username, password, email string) error
	GetUserInfo(ctx context.Context, id uint) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, username, password, email string) error {
	// Check if user exists
	if _, err := s.repo.GetByUsername(ctx, username); err == nil {
		return errors.New("user already exists")
	}

	user := &model.User{
		Username: username,
		Password: password, // In a real app, hash the password
		Email:    email,
	}

	return s.repo.Create(ctx, user)
}

func (s *userService) GetUserInfo(ctx context.Context, id uint) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}
