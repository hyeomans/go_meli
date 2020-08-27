package meli

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type categoriesEndpoints struct {
	c Config
}

func newCategoriesEndpoints(c Config) *categoriesEndpoints {
	return &categoriesEndpoints{
		c: c,
	}
}

func (e *categoriesEndpoints) Get(ctx context.Context, categoryID string) (*CategoryDetails, error) {
	baseURL := fmt.Sprintf("%s/categories/%s", e.c.BaseURL, categoryID)

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

	jsonResponse := CategoryDetails{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, ClientError{
			isTemporary: false,
			message:     fmt.Sprintf("could not unmarshal json into struct: %s - raw body: %v", err, string(body)),
			statusCode:  500,
		}
	}

	return &jsonResponse, nil
}
