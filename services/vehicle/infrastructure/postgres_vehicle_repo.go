package infrastructure

import (
	"context"

	"github.com/alechekz/online-car-auction/services/vehicle/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresVehicleRepo is a PostgreSQL implementation of VehicleRepository interface
type PostgresVehicleRepo struct {
	db *pgxpool.Pool
}

// NewPostgresVehicleRepo creates a new instance of PostgresVehicleRepo
func NewPostgresVehicleRepo(conn string) (*PostgresVehicleRepo, error) {
	pool, err := pgxpool.New(context.Background(), conn)
	if err != nil {
		return nil, err
	}
	return &PostgresVehicleRepo{db: pool}, nil
}

// Save saves a vehicle to the PostgreSQL database
func (r *PostgresVehicleRepo) Save(v *domain.Vehicle) error {
	_, err := r.db.Exec(context.Background(),
		`INSERT INTO vehicles (vin, year, msrp, odometer) VALUES ($1, $2, $3, $4)`,
		v.VIN, v.Year, v.MSRP, v.Odometer,
	)
	return err
}

// FindByVIN retrieves a vehicle by its VIN
func (r *PostgresVehicleRepo) FindByVIN(vin string) (*domain.Vehicle, error) {
	row := r.db.QueryRow(context.Background(),
		`SELECT vin, year, msrp, odometer FROM vehicles WHERE vin=$1`, vin,
	)
	var v domain.Vehicle
	if err := row.Scan(&v.VIN, &v.Year, &v.MSRP, &v.Odometer); err != nil {
		return nil, err
	}
	return &v, nil
}

// Update updates an existing vehicle in the PostgreSQL database
func (r *PostgresVehicleRepo) Update(v *domain.Vehicle) error {
	_, err := r.db.Exec(context.Background(),
		`UPDATE vehicles SET year=$1, odometer=$2 WHERE vin=$3`,
		v.Year, v.Odometer, v.VIN,
	)
	return err
}

// Delete removes a vehicle from the PostgreSQL database by its VIN
func (r *PostgresVehicleRepo) Delete(vin string) error {
	_, err := r.db.Exec(context.Background(),
		`DELETE FROM vehicles WHERE vin=$1`, vin,
	)
	return err
}

// List lists all vehicles in the PostgreSQL database
func (r *PostgresVehicleRepo) List() ([]*domain.Vehicle, error) {
	rows, err := r.db.Query(context.Background(),
		`SELECT vin, year, msrp, odometer FROM vehicles`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []*domain.Vehicle
	for rows.Next() {
		var v domain.Vehicle
		if err := rows.Scan(&v.VIN, &v.Year, &v.MSRP, &v.Odometer); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, &v)
	}
	return vehicles, nil
}
