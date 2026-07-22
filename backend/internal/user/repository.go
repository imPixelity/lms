package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *repo {
	return &repo{db: db}
}

func (r *repo) findById(ctx context.Context, id int) (user, error) {
	var u user
	sqlStr := "SELECT id, email, username FROM users WHERE id = $1"
	err := r.db.QueryRow(ctx, sqlStr, id).Scan(&u.id, &u.email, &u.username)
	if err != nil {
		return user{}, fmt.Errorf("TODO %w", err)
	}
	return u, nil
}
