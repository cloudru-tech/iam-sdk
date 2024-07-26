package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const discoveryURL = "https://api.cloud.ru/endpoints"

// EndpointsResponse is a response from the Cloud.ru API.
type EndpointsResponse struct {
	// Endpoints contains the list of actual API addresses of Cloud.ru products.
	Endpoints []Endpoint `json:"endpoints"`
}

// Endpoint is a product API address.
type Endpoint struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

// DiscoverEndpoint retrieves the IAM gRPC server endpoint from the discovery service.
func DiscoverEndpoint() (string, error) {
	req, err := http.NewRequest(http.MethodGet, discoveryURL, http.NoBody)
	if err != nil {
		return "", fmt.Errorf("construct HTTP request for cloud.ru endpoints: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("get cloud.ru endpoints: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get cloud.ru endpoints: unexpected status code %d", resp.StatusCode)
	}

	var er EndpointsResponse
	if err = json.NewDecoder(resp.Body).Decode(&er); err != nil {
		return "", fmt.Errorf("decode cloud.ru endpoints: %w", err)
	}

	for i := range er.Endpoints {
		if er.Endpoints[i].ID == "iam" {
			return er.Endpoints[i].Address, nil
		}
	}

	return "", fmt.Errorf("iam API is not available")
}
