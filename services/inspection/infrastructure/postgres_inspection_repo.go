package infrastructure

import (
	"context"

	"github.com/alechekz/online-car-auction/services/inspection/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresInspectionRepo is a PostgreSQL implementation of InspectionRepository interface
type PostgresInspectionRepo struct {
	db *pgxpool.Pool
}

// NewPostgresInspectionRepo creates a new instance of PostgresInspectionRepo
func NewPostgresInspectionRepo(conn string) (*PostgresInspectionRepo, error) {
	pool, err := pgxpool.New(context.Background(), conn)
	if err != nil {
		return nil, err
	}
	return &PostgresInspectionRepo{db: pool}, nil
}

// Save saves a vehicle to the PostgreSQL database
func (r *PostgresInspectionRepo) Save(v *domain.Inspection) error {
	_, err := r.db.Exec(context.Background(),
		`INSERT INTO inspections (vin, year) VALUES ($1, $2)`,
		v.VIN, v.Year,
	)
	return err
}

// FindByVIN retrieves a vehicle by its VIN
func (r *PostgresInspectionRepo) FindByVIN(vin string) (*domain.Inspection, error) {
	row := r.db.QueryRow(context.Background(),
		`SELECT vin, year FROM inspections WHERE vin=$1`, vin,
	)
	var v domain.Inspection
	if err := row.Scan(&v.VIN, &v.Year); err != nil {
		return nil, err
	}
	return &v, nil
}
