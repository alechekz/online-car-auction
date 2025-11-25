package domain

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
)

// Vehicle represents a vehicle entity in the system
type Vehicle struct {
	VIN             string `json:"vin"`
	Year            int32  `json:"year"`
	Odometer        int32  `json:"odometer"`
	ExteriorColor   string `json:"exteriorColor"`
	InteriorColor   string `json:"interiorColor"`
	MSRP            uint64 `json:"msrp"`
	Price           uint64 `json:"price"`
	Grade           int    `json:"grade"`
	SmallScratches  bool   `json:"small_scratches"`
	StrongScratches bool   `json:"strong_scratches"`
	ElectricFail    bool   `json:"electric_fail"`
	SuspensionFail  bool   `json:"suspension_fail"`
	Brand           string `json:"brand"`
	Engine          string `json:"engine"`
	Transmission    string `json:"transmission"`
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
			&v.Odometer,
			validation.Min(0),
		),
	)
}
