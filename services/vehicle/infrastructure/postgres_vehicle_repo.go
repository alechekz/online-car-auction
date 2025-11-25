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
		`INSERT INTO vehicles
		(vin, year, odometer, brand, engine, transmission, msrp, grade, price, exterior_color, interior_color, small_scratches, strong_scratches, electric_fail, suspension_fail)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`,
		v.VIN, v.Year, v.Odometer, v.Brand, v.Engine, v.Transmission, v.MSRP, v.Grade, v.Price,
		v.ExteriorColor, v.InteriorColor, v.SmallScratches, v.StrongScratches, v.ElectricFail, v.SuspensionFail,
	)
	return err
}

// FindByVIN retrieves a vehicle by its VIN
func (r *PostgresVehicleRepo) FindByVIN(vin string) (*domain.Vehicle, error) {
	row := r.db.QueryRow(context.Background(),
		`SELECT vin, year, msrp, odometer, brand, engine, transmission, grade, price, exterior_color, interior_color, small_scratches, strong_scratches, electric_fail, suspension_fail
		FROM vehicles WHERE vin=$1`, vin,
	)
	var v domain.Vehicle
	if err := row.Scan(
		&v.VIN, &v.Year, &v.MSRP, &v.Odometer, &v.Brand, &v.Engine, &v.Transmission, &v.Grade, &v.Price, &v.ExteriorColor, &v.InteriorColor, &v.SmallScratches, &v.StrongScratches, &v.ElectricFail, &v.SuspensionFail,
	); err != nil {
		return nil, err
	}
	return &v, nil
}

// Update updates an existing vehicle in the PostgreSQL database
func (r *PostgresVehicleRepo) Update(v *domain.Vehicle) error {
	_, err := r.db.Exec(context.Background(),
		`UPDATE vehicles
		SET year=$1, odometer=$2, brand=$3, engine=$4, transmission=$5, msrp=$6, grade=$7, price=$8, exterior_color=$9, interior_color=$10, small_scratches=$11, strong_scratches=$12, electric_fail=$13, suspension_fail=$14
		WHERE vin=$15`,
		v.Year, v.Odometer, v.Brand, v.Engine, v.Transmission, v.MSRP, v.Grade, v.Price, v.ExteriorColor, v.InteriorColor, v.SmallScratches, v.StrongScratches, v.ElectricFail, v.SuspensionFail, v.VIN,
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
		`SELECT
		vin, year, msrp, odometer, brand, engine, transmission, grade, price, exterior_color, interior_color, small_scratches, strong_scratches, electric_fail, suspension_fail
		FROM vehicles`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []*domain.Vehicle
	for rows.Next() {
		var v domain.Vehicle
		if err := rows.Scan(
			&v.VIN, &v.Year, &v.MSRP, &v.Odometer, &v.Brand, &v.Engine, &v.Transmission, &v.Grade, &v.Price, &v.ExteriorColor, &v.InteriorColor, &v.SmallScratches, &v.StrongScratches, &v.ElectricFail, &v.SuspensionFail,
		); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, &v)
	}
	return vehicles, nil
}
