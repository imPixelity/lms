package user

import "context"

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

type Service interface {
	Get(ctx context.Context, userId int) (*User, error)
}

func (s *service) Get(ctx context.Context, userId int) (*User, error) {
	return s.repo.FindByID(ctx, userId)
}
