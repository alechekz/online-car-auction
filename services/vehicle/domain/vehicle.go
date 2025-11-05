package domain

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
)

// Vehicle represents a vehicle entity in the system
type Vehicle struct {
	VIN           string  `json:"vin"`
	Year          int     `json:"year"`
	Odometer      int     `json:"odometer"`
	ExteriorColor string  `json:"exteriorColor"`
	InteriorColor string  `json:"interiorColor"`
	MSRP          float64 `json:"msrp"`
}

// Validate checks if the vehicle data is valid
func (v *Vehicle) Validate() error {
	return validation.ValidateStruct(
		v,
		validation.Field(
			&v.VIN,
			validation.Required,
			validation.Length(17, 17),
		),
		validation.Field(
			&v.Year,
			validation.Required,
			validation.Min(1900),
			validation.Max(time.Now().Year()),
		),
		validation.Field(
			&v.MSRP,
			validation.Required,
			validation.Min(0.0),
		),
		validation.Field(
			&v.Odometer,
			validation.Min(0),
		),
	)
}
