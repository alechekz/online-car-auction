package domain

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

// BuildData represents the build data of a vehicle
type BuildData struct {
	VIN          string `json:"vin"`
	Brand        string `json:"brand"`
	Engine       string `json:"engine"`
	Transmission string `json:"transmission"`
}

// Validate checks if the VIN is valid
func (v *BuildData) Validate() error {
	return validation.ValidateStruct(
		v,
		validation.Field(
			&v.VIN,
			validation.Required,
			validation.Length(17, 17),
		),
	)
}
