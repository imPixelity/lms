package user

import (
	"context"
	"fmt"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

type Service interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, userID int64) (*User, error)
	// List(ctx context.Context, cursor, limit int) ([]User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID int64) error
}

func (s *service) Create(ctx context.Context, user *User) error {
	if err := s.repo.Create(ctx, user); err != nil {
		return fmt.Errorf("TODO %w", err)
	}
	return nil
}

func (s *service) FindByID(ctx context.Context, userID int64) (*User, error) {
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("TODO %w", err)
	}
	return user, nil
}

func (s *service) Update(ctx context.Context, user *User) error {
	if err := s.repo.Update(ctx, user); err != nil {
		return fmt.Errorf("TODO %w", err)
	}
	return nil
}

func (s *service) Delete(ctx context.Context, userID int64) error {
	if err := s.repo.Delete(ctx, userID); err != nil {
		return fmt.Errorf("TODO %w", err)
	}
	return nil
}
