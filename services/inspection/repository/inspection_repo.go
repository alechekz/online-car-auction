package repository

import "github.com/alechekz/online-car-auction/services/inspection/domain"

// InspectionRepository defines the interface for inspection data operations
type InspectionRepository interface {
	Save(v *domain.Inspection) error
	FindByVIN(vin string) (*domain.Inspection, error)
}
