package usecase

// InspectionProvider defines the interface for fetching vehicle build data
type InspectionProvider interface {
	GetMsrp(vin string) (uint64, error)
}
