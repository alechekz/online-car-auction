package domain

import validation "github.com/go-ozzo/ozzo-validation/v4"

// VehiclesBulk represents a bulk operation on vehicles
type VehiclesBulk struct {
	Vehicles []*Vehicle `json:"vehicles"`
}

// Validate checks if the VehiclesBulk data is valid
func (vb *VehiclesBulk) Validate() error {
	return validation.ValidateStruct(
		vb,
		validation.Field(
			&vb.Vehicles,
			validation.Required,
			validation.Each(
				validation.Required,
			),
		),
	)
}
