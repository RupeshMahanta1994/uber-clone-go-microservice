package repository

import (
	"context"
	"errors"
	"github/rupeshmahanta/user-service/internal/models"

	"github.com/jackc/pgx/v5"
)

type IProfileRepository interface {
	CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	GetProfileByUserId(ctx context.Context, userId string) (*models.Profile, error)
	UpdateProfile(ctx context.Context, profile *models.Profile) error
}

type ProfileRepository struct {
	db *pgx.Conn
}

func NewProfileRepository(db *pgx.Conn) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	query := `INSERT INTO profiles (user_id,name,phone,rating) VALUES ($1,$2,$3,$4)`
	_, err := r.db.Exec(ctx, query, profile.UserID, profile.Name, profile.Phone, profile.Rating)
	return profile, err
}

func (r *ProfileRepository) GetProfileByUserId(ctx context.Context, userId string) (*models.Profile, error) {
	query := `SELECT user_id,name,phone,rating,created_at FROM profiles WHERE user_id=$1`
	row := r.db.QueryRow(ctx, query, userId)
	var profile models.Profile
	err := row.Scan(
		&profile.UserID,
		&profile.Name,
		&profile.Phone,
		&profile.Rating,
		&profile.CreatedAt,
	)
	if err != nil {
		return nil, errors.New("Profile not found")
	}
	return &profile, nil
}

func (r *ProfileRepository) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	query := `UPDATE profiles SET name=$1, phone=$2, rating=$3 WHERE user_id=$4`
	rows, err := r.db.Exec(ctx, query, profile.Name, profile.Phone, profile.Rating, profile.UserID)
	if err != nil {
		return errors.New("Error in updating profile data in DB")
	}
	if rows.RowsAffected() == 0 {
		return errors.New("Profile not found or no changes made")
	}
	return err
}
