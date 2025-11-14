package usecase

import "github.com/alechekz/online-car-auction/services/inspection/domain"

type BuildDataProvider interface {
	Fetch(*domain.BuildData) error
}
