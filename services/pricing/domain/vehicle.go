package domain

import (
	"math"
	"strings"

	"github.com/go-ozzo/ozzo-validation/v4"
)

// Vehicle represents a vehicle entity in the system
type Vehicle struct {
	VIN           string `json:"vin"`
	Odometer      int    `json:"odometer"`
	Grade         int    `json:"grade"`
	ExteriorColor string `json:"exterior_color"`
	InteriorColor string `json:"interior_color"`
	Msrp          uint64 `json:"msrp"`
	Price         uint64 `json:"price"`
}

// Validate checks if the data for inspection is valid
func (v *Vehicle) Validate() error {
	return validation.ValidateStruct(
		v,
		validation.Field(
			&v.VIN,
			validation.Required,
			validation.Length(17, 17),
		),
		validation.Field(
			&v.Grade,
			validation.Required,
			validation.Min(1),
			validation.Max(50),
		),
		validation.Field(
			&v.Odometer,
			validation.Min(0),
		),
	)
}

// CalcPrice calculates the recommended price for the vehicle
func (v *Vehicle) CalcPrice() {

	//apply grade factor
	factor := 0.5 + float64(v.Grade)/100.0
	price := float64(v.Msrp) * factor

	//apply odometer factor
	if v.Odometer > 100_000 {
		price = price * 0.9
	} else if v.Odometer > 50_000 {
		price = price * 0.95
	}

	//apply exterior color factors
	switch strings.ToLower(v.ExteriorColor) {
	case "black", "white", "silver":
		price = price * 1.02
	case "red", "yellow", "green":
		price = price * 0.95
	}

	//apply interior color factors
	switch strings.ToLower(v.InteriorColor) {
	case "black", "grey":
		price = price * 1.01
	case "brown", "cream":
		price = price * 1.03
	case "red", "blue", "white":
		price = price * 0.96
	}

	//save final price
	v.Price = uint64(math.Round(price))
}
