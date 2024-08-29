package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas() (LocationAreas, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("Error: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("Error reading response body: %w", err)
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("Error unmarshaling response: %w", err)
	}

	return locationAreas, nil

}