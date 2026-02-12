package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rupeshmahanta/auth-service/internal/model"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByEmailId(ctx context.Context, email string) (*model.User, error)
}

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (id,email,password,role) VALUES ($1,$2,$3,$4)`
	_, error := r.db.Exec(ctx, query, user.ID, user.Email, user.Password, user.Role)
	return error
}
func (r *UserRepository) GetByEmailId(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id,email,password,role FROM users WHERE email=$1`
	row := r.db.QueryRow(ctx, query, email)
	var user model.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
