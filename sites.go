package meli

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type sitesEndpoints struct {
	c Config
}

func newSitesEndpoints(c Config) *sitesEndpoints {
	return &sitesEndpoints{
		c: c,
	}
}

// All retrieve all sites.
func (e *sitesEndpoints) All(ctx context.Context) ([]Site, error) {
	baseURL := fmt.Sprintf("%s/sites", e.c.BaseURL)

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)

	if err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not create new request: %s", err),
			statusCode:  500,
		}
	}

	body, err := request(ctx, req, e.c.HTTP)
	if err != nil {
		return nil, err
	}

	jsonResponse := []Site{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not unmarshal json into struct: %s - raw body: %v", err, string(body)),
			statusCode:  500,
		}
	}

	return jsonResponse, nil
}

func (e *sitesEndpoints) ListingPrices(ctx context.Context, site string, price int) ([]ListingPrice, error) {
	baseURL := fmt.Sprintf("%s/sites/%s/listing_prices?price=%d", e.c.BaseURL, site, price)

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)

	if err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not create new request: %s", err),
			statusCode:  500,
		}
	}

	body, err := request(ctx, req, e.c.HTTP)
	if err != nil {
		return nil, err
	}

	jsonResponse := []ListingPrice{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not unmarshal json into struct: %s - raw body: %v", err, string(body)),
			statusCode:  500,
		}
	}

	return jsonResponse, nil
}

func (e *sitesEndpoints) ListingTypes(ctx context.Context, site string) ([]ListingType, error) {
	baseURL := fmt.Sprintf("%s/sites/%s/listing_types", e.c.BaseURL, site)

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)

	if err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not create new request: %s", err),
			statusCode:  500,
		}
	}

	body, err := request(ctx, req, e.c.HTTP)
	if err != nil {
		return nil, err
	}

	jsonResponse := []ListingType{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not unmarshal json into struct: %s - raw body: %v", err, string(body)),
			statusCode:  500,
		}
	}

	return jsonResponse, nil
}

func (e *sitesEndpoints) Categories(ctx context.Context, site string) ([]Category, error) {
	baseURL := fmt.Sprintf("%s/sites/%s/categories", e.c.BaseURL, site)

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)

	if err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not create new request: %s", err),
			statusCode:  500,
		}
	}

	body, err := request(ctx, req, e.c.HTTP)
	if err != nil {
		return nil, err
	}

	jsonResponse := []Category{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not unmarshal json into struct: %s - raw body: %v", err, string(body)),
			statusCode:  500,
		}
	}

	return jsonResponse, nil
}
