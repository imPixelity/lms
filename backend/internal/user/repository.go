package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, userID int64) (*User, error)
	List(ctx context.Context, cursor int64, limit int) ([]User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID int64) error
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (email, username, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRow(ctx, query, user.Email, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("TODO %w", err)
	}
	return nil
}

func (r *repository) FindByID(ctx context.Context, userID int64) (*User, error) {
	var user User
	query := `
		SELECT id, email, username 
		FROM users 
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, userID).Scan(&user.ID, &user.Email, &user.Username)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("TODO %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("TODO %w", err)
	}
	return &user, nil
}

func (r *repository) Update(ctx context.Context, user *User) error {
	query := `
		UPDATE users
		SET email = $1, username = $2, password = $3
		WHERE id = $4
	`
	tag, err := r.db.Exec(ctx, query, user.Email, user.Username, user.Password, user.ID)
	if err != nil {
		return fmt.Errorf("TODO %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("TODO %w", err)
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, userID int64) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`
	tag, err := r.db.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("TODO %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("TODO %w", err)
	}
	return nil
}

func (r *repository) List(ctx context.Context, cursor int64, limit int) ([]User, error) {
	query := `
		SELECT id, email, username
		FROM users
		WHERE id > $1
		ORDER BY id
		LIMIT $2
	`
	rows, err := r.db.Query(ctx, query, cursor, limit)
	if err != nil {
		return nil, fmt.Errorf("TODO %w", err)
	}
	defer rows.Close()

	users := make([]User, 0, limit)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Username); err != nil {
			return nil, fmt.Errorf("TODO %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("TODO %w", err)
	}

	return users, nil
}
