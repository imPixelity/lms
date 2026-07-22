package user

import "context"

type service struct {
	repo *repo
}

func NewService(repo *repo) *service {
	return &service{repo: repo}
}

func (svc *service) get(ctx context.Context, id int) (user, error) {
	return svc.repo.findById(ctx, id)
}
