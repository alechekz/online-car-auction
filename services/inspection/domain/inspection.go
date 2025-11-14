package domain

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
)

// Inspection represents a inspection entity in the system
type Inspection struct {
	VIN             string `json:"vin"`
	Year            int    `json:"year"`
	Odometer        int    `json:"odometer"`
	Grade           int    `json:"grade"`
	SmallScratches  bool   `json:"small_scratches"`
	StrongScratches bool   `json:"strong_scratches"`
	ElectricFail    bool   `json:"electric_fail"`
	SuspensionFail  bool   `json:"suspension_fail"`
}

// Validate checks if the inspection data is valid
func (v *Inspection) Validate() error {
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

// Inspect calculates and sets the grade of the vehicle based on its condition
func (v *Inspection) Inspect() {

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
	if v.Odometer > 300000 && tempGrade > 30.0 {
		tempGrade = 30.0
	}

	// Save grade
	v.Grade = int(tempGrade)
}
