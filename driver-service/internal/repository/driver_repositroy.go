package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/rupeshmahanta/driver-service/internal/model"
)

type DriverRepository struct {
	db *pgx.Conn
}

func NewDriverRepository(db *pgx.Conn) *DriverRepository {
	return &DriverRepository{db: db}
}
func (r *DriverRepository) Create(ctx context.Context, driver *model.Driver) (string, error) {
	query := `INSERT INTO drivers (user_id,vehicle_number,vehicle_type,status,is_available,rating) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := r.db.Exec(ctx, query, driver.UserId, driver.VechileNubmer, driver.VechileType, driver.Status, driver.IsAvailable, driver.Rating)
	if err != nil {
		log.Println("Error in inserting Driver data", err)
		return "", err
	}
	log.Println("Driver data inserted successfully")
	return "Driver data inserted", nil
}
func (r *DriverRepository) GetByUserId(ctx context.Context, userId string) (*model.Driver, error) {
	query := `SELECT * FROM drivers WHERE user_id=$1`
	row := r.db.QueryRow(ctx, query, userId)
	var driver model.Driver
	err := row.Scan(&driver.UserId, &driver.VechileNubmer, &driver.VechileType, &driver.Status, &driver.IsAvailable, &driver.Rating)
	if err != nil {
		log.Println("Error in getting Driver Info", err)
		return nil, errors.New("Error in getting driver data")
	}
	return &driver, nil
}
func (r *DriverRepository) Update(ctx context.Context, driver *model.Driver) (string, error) {
	query := `UPDATE drivers set vehicle_number=$1,vehicle_type=$2,status=$3,is_available=$4,rating=$5`
	_, err := r.db.Exec(ctx, query, driver.VechileNubmer, driver.VechileType, driver.Status, driver.IsAvailable, driver.Rating)
	if err != nil {
		log.Println("Error in updating driver info", err)
		return "", nil
	}
	return "Driver data updated", nil
}
