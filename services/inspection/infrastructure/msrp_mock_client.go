package infrastructure

import (
	"github.com/alechekz/online-car-auction/services/inspection/domain"
)

// MockMSRPClient is a mock client for fetching MSRP data based on VIN
type MockMSRPClient struct {
	data        map[string]uint64
	defaultMSRP uint64
}

// NewMockMSRPClient creates a new MockMSRPClient instance
func NewMockMSRPClient() *MockMSRPClient {
	return &MockMSRPClient{
		data: map[string]uint64{
			"5YJSA1E26MF168123": 99000, // Tesla Model S
			"WBS8M9C59J5G12345": 72000, // BMW M3
			"1HGCM82633A004352": 25000, // Honda Accord
			"3FA6P0K92HR123456": 28000, // Ford Fusion
			"JHMFA16586S123456": 22000, // Honda Civic Hybrid
		},
		defaultMSRP: 30000, // default MSRP if VIN not found
	}
}

// Fetch fetches the MSRP for a vehicle by its VIN
func (c *MockMSRPClient) Fetch(v *domain.Vehicle) error {
	v.MSRP = c.defaultMSRP
	if msrp, ok := c.data[v.VIN]; ok {
		v.MSRP = msrp
	}
	return nil
}
