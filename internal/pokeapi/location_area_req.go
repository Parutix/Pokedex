package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreas, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	dat, ok := c.cache.Get(fullURL)
	if ok {
		// Cache found
		locationAreas := LocationAreas{}
		err := json.Unmarshal(dat, &locationAreas)
		if err != nil {
			return LocationAreas{}, fmt.Errorf("Error unmarshaling response: %w", err)
		}
		return locationAreas, nil
	} 

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

	c.cache.Add(fullURL, data)
	
	return locationAreas, nil

}

func (c *Client) GetLocationPokemon(locationName string) (Location, error) {
	endpoint := "/location-area/" + locationName
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Location{}, fmt.Errorf("Error: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, fmt.Errorf("Error reading response body: %w", err)
	}

	location := Location{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return Location{}, fmt.Errorf("Error unmarshaling response: %w", err)
	}
	return location, nil
}