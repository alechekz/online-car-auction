package domain

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
)

// Vehicle represents a vehicle entity in the system
type Vehicle struct {
	VIN string `json:"vin"`

	//build data
	Brand        string `json:"brand"`
	Engine       string `json:"engine"`
	Transmission string `json:"transmission"`
	MSRP         uint64 `json:"msrp"`

	//inspection data
	Year            int  `json:"year"`
	Odometer        int  `json:"odometer"`
	Grade           int  `json:"grade"`
	SmallScratches  bool `json:"small_scratches"`
	StrongScratches bool `json:"strong_scratches"`
	ElectricFail    bool `json:"electric_fail"`
	SuspensionFail  bool `json:"suspension_fail"`
}

// Validate checks if the VIN is valid
func (v *Vehicle) ValidateVIN() error {
	return validation.ValidateStruct(
		v,
		validation.Field(
			&v.VIN,
			validation.Required,
			validation.Length(17, 17),
		),
	)
}

// Validate checks if the data for inspection is valid
func (v *Vehicle) Validate() error {
	if err := v.ValidateVIN(); err != nil {
		return err
	}
	return validation.ValidateStruct(
		v,
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

// Inspect calculates and sets the grade of the vehicle based on its condition
func (v *Vehicle) Inspect() {

	//prepare
	tempGrade := 50.0
	curYear := time.Now().Year()

	// Calculate grade
	tempGrade -= float64(curYear - v.Year)
	if v.StrongScratches {
		tempGrade /= 1.08
	}
	if v.SmallScratches {
		tempGrade /= 1.04
	}
	if v.ElectricFail {
		tempGrade /= 1.08
	}
	if v.SuspensionFail {
		tempGrade /= 1.06
	}
	if v.Odometer > 300_000 && tempGrade > 30.0 {
		tempGrade = 30.0
	}

	// Save grade
	v.Grade = int(tempGrade)
}
