package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alechekz/online-car-auction/services/inspection/domain"
)

// NHTSABuildDataClient is a client for fetching build data from the NHTSA API
type NHTSABuildDataClient struct {
	baseURL string
}

// NewNHTSABuildDataClient creates a new NHTSABuildDataClient instance
func NewNHTSABuildDataClient() *NHTSABuildDataClient {
	return &NHTSABuildDataClient{
		baseURL: "https://vpic.nhtsa.dot.gov/api/vehicles",
	}
}

// vpicResponse represents the structure of the NHTSA API response
type vpicResponse struct {
	Results []struct {
		Variable string `json:"Variable"`
		Value    string `json:"Value"`
	} `json:"Results"`
}

// Fetch fetches the build data for a vehicle by its VIN
func (c *NHTSABuildDataClient) Fetch(v *domain.Vehicle) error {

	// Make the HTTP request to NHTSA API
	resp, err := http.Get(
		fmt.Sprintf("%s/DecodeVin/%s?format=json", c.baseURL, v.VIN),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck

	// Decode the response
	var data vpicResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	for _, r := range data.Results {
		switch r.Variable {
		case "Make":
			v.Brand = r.Value
		case "Engine Model":
			v.Engine = r.Value
		case "Transmission Style":
			v.Transmission = r.Value
		}
	}
	return nil
}
